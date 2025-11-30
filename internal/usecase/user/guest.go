package user

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/entity/validation"
)

func (s *service) LoginWtihEmail(email, password string) (*entity.User, error) {
	if validation.ValidateEmail(email) != nil {
		return nil, ErrInvalidEmail
	}

	user, err := s.userStorage.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, ErrInvalidPassword
	}

	return user, nil
}

func (s *service) LoginWtihPhone(phone, password string) (*entity.User, error) {
	if validation.ValidatePhone(phone) != nil {
		return nil, ErrInvalidPhone
	}

	user, err := s.userStorage.GetUserByPhone(phone)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
