package main

import (
	"log"

	"learnbot/src"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const TELEGRAM_API_TOKEN = "6854615778:AAH92L0DhGWvBx_S9T79H1SEMKRYrkkzAJo"
const UPDATE_TIMEOUT = 10
const DEFAULT_RESPONSE = "Hello dear"

func main() {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_API_TOKEN)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized username is %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = UPDATE_TIMEOUT

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		
		resp := src.GetResponse(update.Message)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp)
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
	}
}