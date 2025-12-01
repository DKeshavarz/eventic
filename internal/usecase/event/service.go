package event

import (
	"github.com/DKeshavarz/eventic/internal/entity"
)

func (s *service) Create(event *entity.Event) (*entity.Event, error) {
	if err := event.Validate(); err != nil {
		return nil, err
	}
	newEvent, err := s.eventStorage.Create(event)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}
