package commandHandler

import (
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CmdHelp(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	msg := "**To use this bot, simply follow these steps:**\n\n1. Start a new ZIP file by sending the command /zip to the bot.\n2. Send the bot any files you want to include in the ZIP file, such as images, documents, GIFs, and more.\n3. (Optional) If you want the ZIP file to be password protected, use the command /password YOUR-PASSWORD.\n4. (Optional) If you want to specify a level of compression for the ZIP file, use the command /compress [0-x], where 0 is no compression and x is maximum compression.\n5. When you are finished adding files to the ZIP, use the command /finish to have the bot download and compress the files. The finished ZIP file will be sent back to you."
	helper.SendMarkdownTelegramMessage(bot, m, msg)
}
