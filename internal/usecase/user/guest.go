package user

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/entity/validation"
)

func (g *guest) LoginWtihEmail(email, password string) (*entity.User, error) {
	if validation.ValidateEmail(email) != nil {
		return nil, ErrInvalidEmail
	}

	user, err := g.userStorage.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, ErrInvalidPassword
	}

	return user, nil
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
