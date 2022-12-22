package modules

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
)

func Command(bot *tgbotapi.BotAPI, m *tgbotapi.Message, files map[string]string) {

	if m.IsCommand() {
		switch m.Command() {
		case "zip":
			cmdNotImplemented(*bot, *m)
		case "password":
			cmdNotImplemented(*bot, *m)
		case "finish":
			cmdNotImplemented(*bot, *m)
		case "compression":
			cmdNotImplemented(*bot, *m)
		case "status":
			cmdNotImplemented(*bot, *m)
		default:
			cmdNotImplemented(*bot, *m)
		}
	}

}

func cmdZip(m tgbotapi.Message, files map[string]string) {
	var zipper Zipper
	DownloadFolder, err := TempFolder("./", m.From.UserName)
	zipType := "7zip"

	if err != nil {
		log.Println(err.Error())
	}

	for fileName, fileURL := range files {
		err = DownloadFile(fileURL, DownloadFolder, fileName)
		if err != nil {
			log.Printf(err.Error())
		}

	}

	switch zipType {
	case "zip":
		zipper = NewGoZipper()
	case "7zip":
		zipper = NewSevenZipper()
	default:
		zipper = NewGoZipper()
	}

	zipper.Zip(DownloadFolder, "hallo1234!") //TODO

	defer Clear(DownloadFolder)

}

func cmdPassword(bot tgbotapi.BotAPI, m tgbotapi.Message) string {
	arguments := strings.TrimSpace(m.CommandArguments())
	if arguments != "" {
		SendTelegramMessage(bot, m, "Password set!")
		return arguments
	}
	SendTelegramMessage(bot, m, "No password supplied\n/password [YOUR-PASSWORD]")
	return ""
}

func cmdFinish(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	DownloadFolder, err := TempFolder("./", m.From.UserName)
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

	defer Clear(DownloadFolder + ".zip")
}

func cmdNotImplemented(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	errorMsg := "The command " + m.Text + " is not yet implemented!"
	SendTelegramMessage(bot, m, errorMsg)
}
