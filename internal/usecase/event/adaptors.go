package event

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type Service interface {
	Create(event *entity.Event) (*entity.Event, error)
}

type service struct {
	eventStorage repositories.Event
}

func NewService(eventStorage repositories.Event) Service{
	return &service{
		eventStorage: eventStorage,
	}
}