package repositories

import "github.com/DKeshavarz/eventic/internal/entity"

type User interface {
	GetUserByPhone(phone string) (*entity.User, error)
}

