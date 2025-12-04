package user

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/entity/validation"
)

func (s *service) LoginWithEmail(email, password string) (*entity.User, error) {
	if validation.Email(email) != nil {
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

func (s *service) LoginWithPhone(phone, password string) (*entity.User, error) {
	if validation.Phone(phone) != nil {
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

func (s *service) GetByID(id int) (*entity.User, error) {
	return s.userStorage.GetByID(id)
}
