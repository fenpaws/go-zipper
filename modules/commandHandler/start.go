package commandHandler

import (
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func CmdStart(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	msg := "Hey there! My name is " + strings.Replace(bot.Self.UserName, "_", "", -1) + " and I can help you create ZIP files out of any type of file you send me. Just send me your files and then use the /zip command to create the ZIP.\n\nFor more information about my capabilities, use the /help command."
	helper.SendMarkdownTelegramMessage(bot, m, msg)
}
