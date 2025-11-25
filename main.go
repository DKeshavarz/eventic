package main

import (
	"github.com/DKeshavarz/eventic/internal/config"
	"github.com/DKeshavarz/eventic/internal/delivery"
)

func main() {
	cfg := config.New()

	delivery.Start(cfg.WebServer, cfg.Telegram)
}
