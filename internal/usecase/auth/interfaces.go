package auth

import (
	"errors"
	"time"

	"github.com/DKeshavarz/eventic/internal/getways"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type Service interface {
	SendOTP(mail string, expire time.Duration) error
	VerifyOTP(mail string, code string) error

	getCode(mail string) string
}

var (
	ErrInvalidExpire = errors.New("invalid expire")
	ErrWrongCode     = errors.New("wrong code")
)

type service struct {
	cache  repositories.Cache
	sender getways.Sender
}

func New(cache repositories.Cache, sender getways.Sender) Service {
	return &service{
		cache:  cache,
		sender: sender,
	}
}
