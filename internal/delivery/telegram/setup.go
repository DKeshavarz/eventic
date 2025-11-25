package telegram

import (
	"log"
	"math/rand/v2"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(cfg *Config) error {
	log.Println("\n\n", cfg.APIKey)
	bot, err := tgbotapi.NewBotAPI(cfg.APIKey)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, random())
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

	return  nil
}

var list = []string{
	"شالام",
	"من این سنگ رو دوست ندارم",
	"به حقوق گربه های نارنجی احترام بگذاریم",
	"میو میو",
	"از مگزیک",
	"نه",
	"اره",
	"مالوخ",
	"نه",
	"اره",
	"گل همه رنگش خوبه، بچه زنگش خوبه",
	"من دلخورم",
	"بسه دیگه... برو کار کن",
}
func random() string{
	num := rand.IntN(len(list))
	return list[num]
}
