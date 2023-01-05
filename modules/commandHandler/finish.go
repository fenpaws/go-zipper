package commandHandler

import (
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func CmdFinish(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	DownloadFolder, err := helper.TempFolder("./", m.From.UserName)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	zipBytes, err := os.ReadFile(DownloadFolder + ".zip")
	if err != nil {
		log.Printf(err.Error())
		return
	}

	zipFileBytes := tgbotapi.FileBytes{
		Name:  DownloadFolder + ".zip",
		Bytes: zipBytes,
	}

	m, err = bot.Send(tgbotapi.NewDocument(m.Chat.ID, zipFileBytes))
	if err != nil {
		log.Printf(err.Error())
		return
	}

	defer helper.Clear(DownloadFolder + ".zip")
}
