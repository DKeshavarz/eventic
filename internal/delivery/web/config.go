package web

import (
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
)

type Config struct {
	Port         string `env:"port"`
	Token        *jwt.Config
	RefreshToken *jwt.Config
}

func DefaultConfig() *Config {
	return &Config{
		Port: "8080",
		Token: &jwt.Config{
			Duration: 30 * time.Minute,
			Secret: []byte("Black_Cat"),
		},
		RefreshToken:  &jwt.Config{
			Duration: 24 * time.Hour,
			Secret: []byte("Orange_Cat"),
		},
	}
}
