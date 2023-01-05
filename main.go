package main

import (
	"fmt"
	"github.com/fenpaws/go-zipper/modules/commandHandler"
	"github.com/fenpaws/go-zipper/modules/config"
	"github.com/fenpaws/go-zipper/modules/errors"
	"github.com/fenpaws/go-zipper/modules/fileHandler"
	"github.com/fenpaws/go-zipper/modules/helper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	Cfg, err := config.New()
	if err != nil {
		log.Println("No configuration found, please provide a .env file or set it in your host")
		log.Fatalf(err.Error())
	}

	// Initialize Telegram bot API
	bot, err := tgbotapi.NewBotAPI(Cfg.BotToken)
	if err != nil {
		log.Fatalf("Error creating Telegram bot API: %v", err)
	}
	bot.SetAPIEndpoint(Cfg.TelegramAPIServer + "/bot%s/%s")
	bot.Debug = Cfg.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set update configuration
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// Get updates channel and iterate over updates
	updatesChan := bot.GetUpdatesChan(updateConfig)
	for update := range updatesChan {
		// Process update if it contains a message
		if update.Message != nil {

			if update.Message.Photo != nil {
				// Add photo to files map if present in update message
				err, name, url := fileHandler.AddPhotoToFiles(bot, update)
				if err != nil {
					log.Printf("Error adding photo to files map: %v", err)
					errors.FileToBig(err, bot, update)
				}

				//TESTING #TODO: Remove
				message := fmt.Sprintf("Filename: %s\n;URL: %s", name, url)
				helper.SendTelegramMessage(*bot, *update.Message, message)
			}

			if update.Message.Document != nil || update.Message.Sticker != nil || update.Message.Voice != nil {
				// Add document to files map if present in update message
				err, name, url := fileHandler.AddFileToFiles(bot, update)
				if err != nil {
					log.Printf("Error adding file to files map: %v", err)
					errors.FileToBig(err, bot, update)
				}

				//TESTING #TODO: Remove
				message := fmt.Sprintf("Filename: %s\nURL: %s", name, url)
				helper.SendTelegramMessage(*bot, *update.Message, message)
			}

			// Handle command if present in update message
			commandHandler.Command(bot, update.Message)

			// Print full JSON message if debugging is enabled
			if Cfg.Debug {
				if !update.Message.IsCommand() {
					log.Printf(update.Message.Text)
				}
			}

		}
	}
}
