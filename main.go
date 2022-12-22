package main

import (
	"fmt"
	"github.com/fenpaws/go-zipper/modules/commands"
	"github.com/fenpaws/go-zipper/modules/config"
	"github.com/fenpaws/go-zipper/modules/helper"
	"github.com/fenpaws/go-zipper/modules/telegramfiles"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	Cfg, err := config.New()

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

	// Initialize map to store files
	files := make(map[string]string)

	// Get updates channel and iterate over updates
	updatesChan := bot.GetUpdatesChan(updateConfig)
	for update := range updatesChan {
		// Process update if it contains a message
		if update.Message != nil {
			// Add document to files map if present in update message
			if err := telegramfiles.AddFileToFiles(bot, update, files); err != nil {
				log.Printf("Error adding file to files map: %v", err)
				helper.ErrorFileToBig(err, bot, update)

			}

			// Add photo to files map if present in update message
			if err := telegramfiles.AddPhotoToFiles(bot, update, files); err != nil {
				log.Printf("Error adding photo to files map: %v", err)
				helper.ErrorFileToBig(err, bot, update)
			}

			// Handle command if present in update message
			commands.Command(bot, update.Message, files)

			// Print full JSON message if debugging is enabled
			if Cfg.Debug {
				if !update.Message.IsCommand() {
					log.Printf(update.Message.Text)
					for key, value := range files {
						fmt.Printf("key: %s, value: %s", key, value)
					}
				}
			}

			// Clear files map
			files = make(map[string]string)
		}
	}
}
