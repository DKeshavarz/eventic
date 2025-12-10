package mail

import (
	"github.com/DKeshavarz/eventic/internal/getways"
	gomail "gopkg.in/mail.v2"
)

type sender struct {
	From string
	Key  string
}

type Config struct {
	From string
	Key  string
}

func New(cfg *Config) getways.Sender {
	return &sender{
		From: cfg.From,
		Key:  cfg.Key,
	}
}

func (s *sender) Send(to string, message *getways.Message) error {
	msg := gomail.NewMessage()

	msg.SetHeader("From", s.From)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", message.Title)
	msg.SetBody("text/plain", message.Text)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, s.From, s.Key)

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
