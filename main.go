package main

import (
	"github.com/DKeshavarz/eventic/internal/config"
	"github.com/DKeshavarz/eventic/internal/delivery"
	"github.com/DKeshavarz/eventic/internal/getways/mail"
	"github.com/DKeshavarz/eventic/internal/repositories/cache"
	"github.com/DKeshavarz/eventic/internal/repositories/inmemory"
	"github.com/DKeshavarz/eventic/internal/usecase/auth"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
)

func main() {
	cfg := config.New()

	db := inmemory.DefaultDB()
	cache := cache.New()
	sender := mail.New(cfg.Mail)
	userStorage := inmemory.NewUserStorage(db)

	// orgStorage := inmemory.NewOrgStorage(db)
	// eventStorage := inmemory.NewEventStorage(db)
	// joinEventStorage := inmemory.NewJoinEventStorage(db)

	userSevice := user.NewSevice(userStorage)
	authService := auth.New(cache, sender)
	delivery.Start(cfg.Delivery, userSevice, authService)
}
