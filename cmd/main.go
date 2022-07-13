package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/storyteller23/telegram-aliyacenter/pkg/config"
	"github.com/storyteller23/telegram-aliyacenter/pkg/controller"
	"github.com/storyteller23/telegram-aliyacenter/pkg/handler"
)

func main() {
	defer catchPanic()
	cfg := config.NewConfig()
	err := controller.CreateTable(cfg.DBName)
	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	fmt.Printf("%T\n", bot)
	for update := range updates {
		go handler.ProcessRequest(update, bot, cfg.DBName)
	}
}

func catchPanic() {
	recover()
	fmt.Println("PANIC")
	fmt.Println("Reload....")
	main()
}
