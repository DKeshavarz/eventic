package main

import (
	"github.com/DKeshavarz/eventic/internal/config"
	"github.com/DKeshavarz/eventic/internal/delivery"
	"github.com/DKeshavarz/eventic/internal/repositories/inmemory"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
)

func main() {
	cfg := config.New()

	db := inmemory.DefaultDB()

	userStorage := inmemory.NewUserStorage(db)
	// orgStorage := inmemory.NewOrgStorage(db)
	// eventStorage := inmemory.NewEventStorage(db)
	// joinEventStorage := inmemory.NewJoinEventStorage(db)

	userSevice := user.NewSevice(userStorage)

	delivery.Start(cfg.WebServer, cfg.Telegram, userSevice)
}
