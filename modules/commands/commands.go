package commands

import (
	"fmt"
	"github.com/fenpaws/go-zipper/modules/downloader"
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/fenpaws/go-zipper/modules/zip"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
)

func Command(bot *tgbotapi.BotAPI, m *tgbotapi.Message, files map[string]string) {

	if m.IsCommand() {
		switch m.Command() {
		case "start":
			cmdStart(*bot, *m)
		case "help":
			cmdHelp(*bot, *m)
		case "about":
			cmdAbout(*bot, *m)
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
func cmdStart(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	msg := "Hey there! My name is " + strings.Replace(bot.Self.UserName, "_", "", -1) + " and I can help you create ZIP files out of any type of file you send me. Just send me your files and then use the /zip command to create the ZIP.\n\nFor more information about my capabilities, use the /help command."
	helper.SendTelegramMessage(bot, m, msg)
}

func cmdHelp(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	msg := "**To use this bot, simply follow these steps:**\n\n1. Start a new ZIP file by sending the command /zip to the bot.\n2. Send the bot any files you want to include in the ZIP file, such as images, documents, GIFs, and more.\n3. (Optional) If you want the ZIP file to be password protected, use the command /password YOUR-PASSWORD.\n4. (Optional) If you want to specify a level of compression for the ZIP file, use the command /compress [0-x], where 0 is no compression and x is maximum compression.\n5. When you are finished adding files to the ZIP, use the command /finish to have the bot download and compress the files. The finished ZIP file will be sent back to you."
	helper.SendTelegramMessage(bot, m, msg)
}

func cmdAbout(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	msg := "Developed by [SoXX](https://t.me/Fenpaws)\nSource Code @ [GitHub](https://github.com/fenpaws/go-zipper)\n"
	helper.SendTelegramMessage(bot, m, msg)
}

func cmdZip(m tgbotapi.Message, files map[string]string) {
	var zipper zip.Zipper
	DownloadFolder, err := helper.TempFolder("./", m.From.UserName)
	zipType := "7zip"

	if err != nil {
		log.Println(err.Error())
	}

	for fileName, fileURL := range files {
		err = downloader.DownloadFile(fileURL, DownloadFolder, fileName)
		if err != nil {
			log.Printf(err.Error())
		}

	}

	switch zipType {
	case "zip":
		zipper = zip.NewGoZipper()
	case "7zip":
		zipper = zip.NewSevenZipper()
	default:
		zipper = zip.NewGoZipper()
	}

	zipper.Zip(DownloadFolder, "hallo1234!") //TODO

	defer helper.Clear(DownloadFolder)

}

func cmdPassword(bot tgbotapi.BotAPI, m tgbotapi.Message) string {
	arguments := strings.TrimSpace(m.CommandArguments())
	if arguments != "" {
		helper.SendTelegramMessage(bot, m, "Password set!")
		return arguments
	}
	helper.SendTelegramMessage(bot, m, "No password supplied\n/password [YOUR-PASSWORD]")
	return ""
}

func cmdFinish(bot tgbotapi.BotAPI, m tgbotapi.Message) {
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

func cmdNotImplemented(bot tgbotapi.BotAPI, m tgbotapi.Message) {
	errorMsg := fmt.Sprintf("The command %s is not yet implemeted!", m.Text)
	helper.SendTelegramMessage(bot, m, errorMsg)
}
