package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type ConfigDatabase struct {
	BotToken          string `env:"BOT_TOKEN"`
	TelegramAPIServer string `env:"TELEGRAM_API_SERVER" envDefault:"https://api.telegram.org"`
	Debug             bool   `env:"DEBUG" envDefault:"false"`
}

func New() (*ConfigDatabase, error) {
	cfg := &ConfigDatabase{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	log.Printf("Config Loaded!")
	log.Printf("Debug enabled: %t", cfg.Debug)
	return cfg, nil
}
