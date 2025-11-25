package config

import (
	"github.com/DKeshavarz/eventic/internal/delivery/telegram"
	"github.com/DKeshavarz/eventic/internal/delivery/web"
)

type Config struct {
	Telegram  *telegram.Config
	WebServer *web.Config
}

func New() *Config {
	config := &Config{
		Telegram:  telegram.DefaultConfig(),
		WebServer: web.DefaultConfig(),
	}

	loadTelegram(config.Telegram)
	LoadWebServer(config.WebServer)
	return config
}

func loadTelegram(cfg *telegram.Config) {
	cfg.APIKey = getEnv("TELEGRAM_API_KEY", "")
}

func LoadWebServer(cfg *web.Config) {
	cfg.Port = getEnv("WEB_PORT", "8080")
}
