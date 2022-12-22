package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/fenpaws/go-zipper/modules"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ConfigDatabase struct {
	BotToken          string `env:"BOT_TOKEN"`
	TelegramAPIServer string `env:"TELEGRAM_API_SERVER" envDefault:"https://api.telegram.org"`
	Debug             bool   `env:"DEBUG" envDefault:"false"`
}

func main() {
	// Parse environment variables
	var cfg ConfigDatabase
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}

	// Initialize Telegram bot API
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalf("Error creating Telegram bot API: %v", err)
	}
	bot.SetAPIEndpoint(cfg.TelegramAPIServer + "/bot%s/%s")
	bot.Debug = cfg.Debug
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
			if err := modules.AddFileToFiles(bot, update, files); err != nil {
				log.Printf("Error adding file to files map: %v", err)
			}

			// Add photo to files map if present in update message
			if err := modules.AddPhotoToFiles(bot, update, files); err != nil {
				log.Printf("Error adding photo to files map: %v", err)
			}

			// Handle command if present in update message
			modules.Command(bot, update.Message, files)

			// Print full JSON message if debugging is enabled
			if cfg.Debug {
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
