package organization

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type Service interface {
	Create(org *entity.Organization) (*entity.Organization, error)
}

var (
	ErrInvalidOwner = errors.New("invalid owner")
	ErrDuplicatedName = errors.New("duplicated name")
)

type service struct {
	orgStorage repositories.Organization
}

func NewService(orgStorage repositories.Organization) Service {
	return &service{
		orgStorage: orgStorage,
	}
}
