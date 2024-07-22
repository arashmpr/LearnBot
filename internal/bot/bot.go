package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const DEFAULT_NON_COMMAND_RESPONSE = "گزینه موجود نیست. لطفا دوباره انتخاب کنید."

const ASK_FOR_PERSONAL_INFO_TEXT = "لطفا اطلاعات زیر را وارد نمایید"
const ASK_FOR_FIRST_NAME_TEXT = "نام"
const ASK_FOR_LAST_NAME_TEXT = "نام خانوادگی"
const ASK_FOR_EMAIL_TEXT = "ایمیل"
const ASK_FOR_PHONE_NO_TEXT = "شماره همراه"
const ASK_FOR_TELEGRAM_ID_TEXT = "آیدی تلگرام"

func HandleCommand(tgbot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	if msg.IsCommand() {
		cmd := msg.Command()
		switch cmd {
			case "start":
				StartHandler(tgbot, msg)
		}
	} else {
		DefaultNonCommandHandler(tgbot, msg)
	}
}


func StartHandler(tgbot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, ASK_FOR_PERSONAL_INFO_TEXT)
	_, err := tgbot.Send(reply)
	if err != nil {
		log.Printf("[Error] Cannot ask for personal info: %v", err)
	}
}

func DefaultNonCommandHandler(tgbot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, DEFAULT_NON_COMMAND_RESPONSE)
	_, err := tgbot.Send(reply)
	if err != nil {
		log.Printf("[Error] non command handler does not send: %v", err)
	}
}


