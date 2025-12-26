package event

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type Service interface {
	Create(event *entity.Event) (*entity.Event, error)
	Join(joinEvent *entity.JoinEvent) (*entity.JoinEvent, error)
	GetAll()([]*entity.Event, error)
}

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidEvent = errors.New("invalid event")
)

type service struct {
	eventStorage     repositories.Event
	joinEventStorage repositories.JoinEvent
}

func NewService(eventStorage repositories.Event, joinEventStorage repositories.JoinEvent) Service {
	return &service{
		eventStorage:     eventStorage,
		joinEventStorage: joinEventStorage,
	}
}
