package main

import (
	"log"

	"learnbot/internal/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const TELEGRAM_API_TOKEN = "6854615778:AAH92L0DhGWvBx_S9T79H1SEMKRYrkkzAJo"
const UPDATE_TIMEOUT = 10
const DEFAULT_RESPONSE = "Hello dear"

func main() {
	tgbot, err := tgbotapi.NewBotAPI(TELEGRAM_API_TOKEN)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	tgbot.Debug = true
	log.Printf("Authorized username is %s", tgbot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = UPDATE_TIMEOUT

	updates, err := tgbot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		
		bot.HandleCommand(tgbot, update.Message)
	}
}