package config

import (
	"fmt"
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery/telegram"
	"github.com/DKeshavarz/eventic/internal/delivery/web"
)

type Config struct {
	Telegram  *telegram.Config
	WebServer *web.Config
}

func New() *Config {
	Load(".env")
	config := &Config{
		Telegram:  telegram.DefaultConfig(),
		WebServer: web.DefaultConfig(),
	}

	loadTelegram(config.Telegram)
	LoadWebServer(config.WebServer)
	fmt.Println(config)
	return config
}

func loadTelegram(cfg *telegram.Config) {
	cfg.APIKey = getEnv("TELEGRAM_API_KEY", "")
}

func LoadWebServer(cfg *web.Config) {
	cfg.Port = getEnv("WEB_PORT", cfg.Port)

	cfg.Token.Secret = []byte(getEnv("JWT_TOKEN_SECRET", string(cfg.Token.Secret[:])))
	duration := getEnvAsInt("JWT_TOKEN_DURATION", int(cfg.Token.Duration))
	cfg.Token.Duration = time.Duration(duration) * time.Hour

	duration = getEnvAsInt("JWT_REFRESH_TOKEN_DURATION", int(cfg.RefreshToken.Duration))
	cfg.RefreshToken.Duration = time.Duration(duration) * time.Hour
	cfg.RefreshToken.Secret = []byte(getEnv("JWT_REFRESH_TOKEN_SECRET", string(cfg.RefreshToken.Secret[:])))
}
