package commands

import (
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CmdAbout(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	msg := "Developed by [SoXX](https://t.me/Fenpaws)\nSource Code @ [GitHub](https://github.com/fenpaws/go-zipper)\n"
	helper.SendTelegramMessage(bot, m, msg)
}
