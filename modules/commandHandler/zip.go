package commandHandler

import (
	"github.com/fenpaws/go-zipper/modules/downloader"
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/fenpaws/go-zipper/modules/zip"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func CmdZip(m tgbotapi.Message, files map[string]string) {
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
