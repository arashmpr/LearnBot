package src

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const DEFAULT_NON_COMMAND_RESPONSE = "گزینه موجود نیست. لطفا دوباره انتخاب کنید."
func GetResponse(m *tgbotapi.Message) string {
	if m.IsCommand() {
		return HandleCommand(m.Command())
	}
	return DEFAULT_NON_COMMAND_RESPONSE
}
func HandleCommand(cmd string) string {
	switch cmd {
		case "start":
			return "لطفا نام کاربری خود را وارد کنید"
		case "stop":
			return "برو سیکت رو بزن عشقولی"
	}
	return "حرفی ندارم"
}

