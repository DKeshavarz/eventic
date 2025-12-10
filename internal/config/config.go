package config

import (
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery"
	"github.com/DKeshavarz/eventic/internal/delivery/telegram"
	"github.com/DKeshavarz/eventic/internal/delivery/web"
	"github.com/DKeshavarz/eventic/internal/getways/mail"
)

type Config struct {
	Delivery *delivery.Config
	Mail     *mail.Config
}

func New() *Config {
	Load(".env")
	cfg := &Config{
		Delivery: &delivery.Config{
			TelegramCofig: telegram.DefaultConfig(),
			WebConfig:     web.DefaultConfig(),
		},
		Mail: &mail.Config{},
	}

	loadTelegram(cfg.Delivery.TelegramCofig)
	LoadWebServer(cfg.Delivery.WebConfig)
	LoadMail(cfg.Mail)

	return cfg
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

func LoadMail(cfg *mail.Config) {
	cfg.From = getEnv("MAIL_FROM", "")
	cfg.Key = getEnv("MAIL_KEY", "")
}
