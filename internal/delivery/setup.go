package delivery

import (
	"log"

	"github.com/DKeshavarz/eventic/internal/delivery/telegram"
	"github.com/DKeshavarz/eventic/internal/delivery/web"
	"github.com/DKeshavarz/eventic/internal/usecase/auth"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
)

type Config struct {
	WebConfig *web.Config
	TelegramCofig *telegram.Config
}
const (
	INTERFACES_COUNT = 2
)

func Start(cfg *Config, userSevice user.Service, authService auth.Service, eventServic event.Service) error{
	ch := make(chan any)

	go func() {
		err := web.Start(cfg.WebConfig, userSevice, eventServic,authService)
		log.Println("web stpos -> ", err)
		ch <- "Done"
	}()

	go func() {
		err := telegram.Start(cfg.TelegramCofig)
		log.Println("telegram stops ->", err)
		ch <- "Done"
	}()

	for range INTERFACES_COUNT{
		<- ch 
	}

	return  nil
}