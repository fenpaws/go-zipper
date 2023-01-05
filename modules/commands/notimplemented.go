package commands

import (
	"fmt"
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CmdNotImplemented(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	errorMsg := fmt.Sprintf("The command %s is not yet implemeted!", m.Text)
	helper.SendTelegramMessage(bot, m, errorMsg)
}
