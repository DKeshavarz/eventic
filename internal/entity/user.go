package entity

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity/validation"
)

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

var (
	ErrWeakPassword = errors.New("پسورد ضعیف است")
	ErrNotEnoughCredential = errors.New("اطلاعات داده شده کافی نیست")
)
func (u *User) Validate() error {
	if ValidatePassword(u.Password) != nil {
		return ErrWeakPassword
	}

	if u.Email == nil && u.Phone == nil {
		return ErrNotEnoughCredential
	}

	if u.Phone != nil && validation.Phone(*u.Phone) != nil {
		return ErrInvalidPhone
	}

	if u.Email != nil && validation.Email(*u.Email) != nil {
		return ErrInvalidEmail
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrWeakPassword
	}
	return nil
}