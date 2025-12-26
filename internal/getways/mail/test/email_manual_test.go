//go:build manual
package test

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/config"
	"github.com/DKeshavarz/eventic/internal/getways"
	"github.com/DKeshavarz/eventic/internal/getways/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	cfg := config.New()

	mail := mail.New(cfg.Mail)

	err := mail.Send("dankeshavarz1075@gmail.com", &getways.Message{
		Title: "This is Title",
		Text:  "This is my text\nthis is text 2",
	})

	assert.Nil(t, err)
}
