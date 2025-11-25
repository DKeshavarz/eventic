package delivery

import (
	"log"

	"github.com/DKeshavarz/eventic/internal/delivery/telegram"
	"github.com/DKeshavarz/eventic/internal/delivery/web"
)

const (
	INTERFACES_COUNT = 2
)

func Setup(webCfg *web.Config, telegramCfg *telegram.Config){
	ch := make(chan any)

	go func() {
		err := web.Start(webCfg)
		log.Println(err)
		ch <- "Done"
	}()

	go func() {
		err := telegram.Start(telegramCfg)
		log.Println(err)
		ch <- "Done"
	}()

	for range INTERFACES_COUNT{
		<- ch 
	}
}