package config

import (
	"os"

	"github.com/DKeshavarz/eventic/internal/delivery/telegram"
)

type Config struct {
	Telegram *telegram.Config
}

func New() *Config {
	config := &Config{
		Telegram: telegram.DefaultConfig(),
	}

	loadTelegram(config.Telegram)
	return config
}

func loadTelegram(cfg *telegram.Config) {
	cfg.APIKey = os.Getenv("TELEGRAM_API_KEY")
}
