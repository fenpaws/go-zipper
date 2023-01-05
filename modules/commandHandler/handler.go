package commandHandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Command(bot *tgbotapi.BotAPI, m *tgbotapi.Message) {

	if m.IsCommand() {
		switch m.Command() {
		case "start":
			CmdStart(*bot, *m)
		case "help":
			CmdHelp(*bot, *m)
		case "about":
			CmdAbout(*bot, *m)
		case "zip":
			CmdNotImplemented(*bot, *m)
		case "password":
			CmdNotImplemented(*bot, *m)
		case "finish":
			CmdNotImplemented(*bot, *m)
		case "compression":
			CmdNotImplemented(*bot, *m)
		case "status":
			CmdNotImplemented(*bot, *m)
		default:
			CmdNotImplemented(*bot, *m)
		}
	}

}
