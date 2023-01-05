package commandHandler

import (
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func CmdPassword(bot tgbotapi.BotAPI, m tgbotapi.Message) string {
	arguments := strings.TrimSpace(m.CommandArguments())
	if arguments != "" {
		helper.SendMarkdownTelegramMessage(bot, m, "Password set!")
		return arguments
	}
	helper.SendMarkdownTelegramMessage(bot, m, "No password supplied\n/password [YOUR-PASSWORD]")
	return ""
}
