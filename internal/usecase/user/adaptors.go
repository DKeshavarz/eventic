package user

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type Service interface {
	LoginWithEmail(email, password string) (*entity.User, error)
	LoginWithPhone(phone, password string) (*entity.User, error)
	GetByID(id int) (*entity.User, error)
	Signup(*entity.User) (*entity.User, error)
}

var (
	ErrInvalidPhone    = errors.New("invalid phone number")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidEmail    = errors.New("invalid email")
)

type service struct {
	userStorage repositories.User
	orgStorage repositories.Organization
}

func NewSevice(userStorage repositories.User, orgStorage repositories.Organization) Service {
	return &service{
		userStorage: userStorage,
		orgStorage: orgStorage,
	}
}
