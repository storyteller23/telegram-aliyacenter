package config

import (
	"log"
	"os"
)

type Config struct {
	DBName      string
	BotToken    string
	CompanyCode string
	Auth        string
}

func NewConfig() *Config {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("bot token not specified on env(BOT_TOKEN)")
	}
	companyCode := os.Getenv("COMPANY_CODE")
	if companyCode == "" {
		log.Fatal("company code not specified on env(COMPANY_TOKEN)")
	}
	auth := os.Getenv("BASIC_AUTH")
	if auth == "" {
		log.Fatal("basic auth not specified on env(COMPANY_TOKEN)")
	}
	return &Config{
		DBName:   "users.sqlite",
		BotToken: botToken,
	}
}
