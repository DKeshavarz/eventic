package entity

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity/validation"
)

type Organization struct {
	ID          int     `json:"organizer_id"`
	OwnerID     int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	LogoPic     *string `json:"logo_pic"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
}

var (
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidPhone       = errors.New("invalid phone")
)

func (org *Organization) Validate() error {
	if len(org.Description) == 0 {
		return ErrInvalidDescription
	}

	if len(org.Name) == 0 {
		return ErrInvalidName
	}

	if org.Email != nil && validation.Email(*org.Email) != nil {
		return ErrInvalidEmail
	}

	if org.Phone != nil && validation.Phone(*org.Phone) != nil {
		return ErrInvalidPhone
	}

	return nil
}
