package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	KeyboardInfo = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Цены"),
			tgbotapi.NewKeyboardButton("Акции узи"),
			tgbotapi.NewKeyboardButton("Наша команда"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Записаться"),
			tgbotapi.NewKeyboardButton("Контакты"),
		),
	)
	Empty = tgbotapi.NewRemoveKeyboard(true)
	Skip  = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Пропустить"),
		),
	)
	Date = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Любая"),
			tgbotapi.NewKeyboardButton("Сегодня"),
			tgbotapi.NewKeyboardButton("Завтра"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("На этой неделе"),
			tgbotapi.NewKeyboardButton("На следующей неделе"),
		),
	)
	Time = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Любое"),
			tgbotapi.NewKeyboardButton("9:00-11:00"),
			tgbotapi.NewKeyboardButton("11:00-13:00"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("14:00-16:00"),
			tgbotapi.NewKeyboardButton("16:00-18:00"),
		),
	)
	Specialist = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Любой"),
			tgbotapi.NewKeyboardButton("Выбрать из списка"),
		),
	)
	SpecialistList = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Шангараева Алия Аблаевна"),
			tgbotapi.NewKeyboardButton("Амиркулова Анара Сагиндыковна"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Онгарбаева Назгуль Халилуллаевна"),
			tgbotapi.NewKeyboardButton("Каримова Бахыт Жарылгаповна"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Базылова Светлана Бауыржановна"),
			tgbotapi.NewKeyboardButton("Бисембина Эльвира Сериковна"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Салимова Клара Васильевна"),
			tgbotapi.NewKeyboardButton("Муратова Жумазия Бактыгалиевна"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Турешева Акбаян Буриевна"),
			tgbotapi.NewKeyboardButton("Бекетова Жанат Жанузаковна"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Любой"),
		),
	)
)
