package main

import (
	"archive/zip"
	"fmt"
	"github.com/caarlos0/env/v6"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ConfigDatabase struct {
	BotToken          string `env:"BOT_TOKEN"`
	TelegramAPIServer string `env:"TELEGRAM_API_SERVER" envDefault:"https://api.telegram.org"`
	Debug             bool   `env:"DEBUG" envDefault:"false"`
}

func DownloadFile(url string, downloadFolder string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Printf("Downloading file: %s in %s", fileName, downloadFolder)

	fileHandle, err := os.OpenFile(downloadFolder+"/"+fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	_, err = io.Copy(fileHandle, resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Download of file: %s in folder %s finished!", fileName, downloadFolder)
	return nil
}

func CreateZipFromFolder(folderPath string) error {
	file, err := os.Create(folderPath + ".zip")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := zip.NewWriter(file)
	defer writer.Close()

	return filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(folderPath), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err

	})
}

func CreateUniqueTempFolder(folderPath string, identifier string) (string, error) {
	// Create a unique temp download folder
	DownloadDir, err := os.MkdirTemp(folderPath, "*-"+identifier)
	if err != nil {
		return "", err
	}
	return DownloadDir, nil
}

func main() {

	cfg := ConfigDatabase{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	token := cfg.BotToken

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf(err.Error())
	}
	bot.SetAPIEndpoint(cfg.TelegramAPIServer + "/bot%s/%s")

	bot.Debug = cfg.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	files := make(map[string]string)

	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		if update.Message != nil {
			if update.Message.Document != nil {
				fileUrl, err := bot.GetFileDirectURL(update.Message.Document.FileID)
				if err != nil {
					log.Printf(err.Error())
				}
				fileName := update.Message.Document.FileName
				files[fileName] = fileUrl
			}

			if update.Message.Text == "/zip" {

				DownloadFolder, err := CreateUniqueTempFolder("./", update.Message.From.UserName)

				if err != nil {
					log.Println(err.Error())
				}

				for fileName, fileURL := range files {
					err = DownloadFile(fileURL, DownloadFolder, fileName)
					if err != nil {
						log.Printf(err.Error())
					}

				}

				CreateZipFromFolder(DownloadFolder)
				err = os.RemoveAll(DownloadFolder)
				if err != nil {
					log.Printf(err.Error())
				}
				log.Printf("Folder %s deleated!", DownloadFolder)

				// TESTING //
				zipBytes, err := os.ReadFile(DownloadFolder + ".zip")
				if err != nil {
					log.Printf(err.Error())
				}

				zipFileBytes := tgbotapi.FileBytes{
					Name:  DownloadFolder + ".zip",
					Bytes: zipBytes,
				}

				message, err := bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, zipFileBytes))
				log.Printf(message.Text)
				files = make(map[string]string)

				err = os.RemoveAll(DownloadFolder + ".zip")
				if err != nil {
					log.Printf(err.Error())
				}
				log.Printf("ZIP file %s deleated!", DownloadFolder+".zip")

			}

		}
	}

}
