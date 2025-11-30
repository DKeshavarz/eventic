package usecase

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/entity/validation"
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

func (g *guest) LoginWtihEmail(email, password string) (*entity.User, error) {
	return nil, nil
}

func (g *guest) LoginWtihPhone(phone, password string) (*entity.User, error) {
	if validation.ValidatePhone(phone) != nil {
		return nil, ErrInvalidPhone
	}

	user, err := g.userStorage.GetUserByPhone(phone)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
