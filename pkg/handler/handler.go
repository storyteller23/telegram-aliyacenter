package handler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/storyteller23/telegram-aliyacenter/pkg/config"
	"github.com/storyteller23/telegram-aliyacenter/pkg/controller"
	"github.com/storyteller23/telegram-aliyacenter/pkg/keyboard"
	"github.com/storyteller23/telegram-aliyacenter/pkg/static"
)

func ProcessRequest(update tgbotapi.Update, bot *tgbotapi.BotAPI, dbName string) {
	if update.Message == nil {
		return
	}
	chatId := update.Message.Chat.ID
	text := update.Message.Text
	err := controller.AddChatID(dbName, chatId)
	if err != nil {
		log.Fatal(err)
	}
	if text == "/start" {
		controller.SetFlagByID(dbName, chatId, "info")
	}
	flag, err := controller.GetByID(dbName, chatId, "flag")
	if err != nil {
		log.Fatal(err)
	}
	msg := tgbotapi.NewMessage(chatId, "")

	switch flag {
	case "info":
		msg.ReplyMarkup = keyboard.KeyboardInfo
		switch strings.ToLower(text) {
		case "/start":
			msg.Text = static.StartText
		case "записаться":
			msg.ReplyMarkup = keyboard.Empty
			msg.Text = "Введите ваше имя:"
			err := controller.SetFlagByID(dbName, chatId, "name")
			if err != nil {
				log.Fatal(err)
			}
		case "цены":
			msg.Text = static.Price
		default:
			msg.Text = static.Unknown
		}
	case "name":
		controller.SetNameByID(dbName, chatId, text)
		controller.SetFlagByID(dbName, chatId, "phone")
		msg.Text = "Введите ваш номер телефона:"
	case "phone":
		controller.SetPhoneByID(dbName, chatId, text)
		controller.SetFlagByID(dbName, chatId, "email")
		msg.Text = "Введите ваш email(электронную почту):"
		msg.ReplyMarkup = keyboard.Skip
	case "email":
		controller.SetEmailByID(dbName, chatId, text)
		controller.SetFlagByID(dbName, chatId, "specialist")
		msg.Text = "Выберите специалиста:"
		msg.ReplyMarkup = keyboard.Specialist
	case "specialist":
		if strings.ToLower(text) == "выбрать из списка" {
			msg.ReplyMarkup = keyboard.SpecialistList
			msg.Text = "Выбрать:"
			break
		}
		controller.SetSpecialistByID(dbName, chatId, text)
		controller.SetFlagByID(dbName, chatId, "date")
		msg.Text = "Предпочтительная дата приёма:"
		msg.ReplyMarkup = keyboard.Date
	case "date":
		controller.SetDateByID(dbName, chatId, text)
		controller.SetFlagByID(dbName, chatId, "time")
		msg.Text = "Предпочтительное время приёма:"
		msg.ReplyMarkup = keyboard.Time
	case "time":
		controller.SetTimeByID(dbName, chatId, text)
		controller.SetFlagByID(dbName, chatId, "info")
		if _, err := bot.Send(tgbotapi.NewMessage(chatId, "Отправляем заявку...")); err != nil {
			log.Panic(err)
		}
		err = SendForm(chatId, dbName)
		if err != nil {
			log.Fatal(err)
		}
		msg.Text = "Заявка успешно отправлена!"
		msg.ReplyMarkup = keyboard.KeyboardInfo
	default:
		controller.SetFlagByID(dbName, chatId, "info")
		msg.Text = static.StartText
	}
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func SendForm(chatId int64, dbName string) error {
	client := &http.Client{}
	requestData, err := controller.GetRequestData(dbName, chatId)
	if err != nil {
		return err
	}
	data := url.Values{
		"APPOINTMENT_DATE[]": {requestData.Date},
		"APPOINTMENT_TIME[]": {requestData.Time},
		"PATIENT_NAME":       {requestData.Name},
		"CONTACTS":           {requestData.Phone},
		"DESCRIPTION":        {requestData.Specialist},
		"COMPANY_CODE":       {config.NewConfig().CompanyCode},
	}

	req, err := http.NewRequest("POST", static.Url, nil)
	if err != nil {
		return err
	}
	req.PostForm = data
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(config.NewConfig().Auth)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	res.Body.Close()

	return nil
}
