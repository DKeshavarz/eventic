package usecase

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type Guest interface {
	LoginWtihEmail(email, password string) (*entity.User, error)
	LoginWtihPhone(phone, password string) (*entity.User, error)
}

var (
	ErrInvalidPhone    = errors.New("invalid phone number")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)

type guest struct {
	userStorage repositories.User
}

func NewGuest(userStorage repositories.User) Guest {
	return &guest{
		userStorage: userStorage,
	}
}
