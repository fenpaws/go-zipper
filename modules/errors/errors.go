package errors

import (
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func FileToBig(err error, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if strings.Contains(err.Error(), "file is too big") {
		msg := "❌ - Im sorry but i cant download that file, its more then 20MB (API Limitation)"
		helper.SendTelegramMessage(*bot, *update.Message, msg)
	}
}
