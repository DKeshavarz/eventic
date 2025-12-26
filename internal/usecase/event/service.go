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

func (s *service) Join(joinEvent *entity.JoinEvent) (*entity.JoinEvent, error) {
	return s.joinEventStorage.Create(joinEvent)
}

func (s *service) GetAll() (events []*entity.Event, err error) {
	events, err = s.eventStorage.GetAll()
	return
}

func (s *service) Get(id int) (event *entity.Event, err error) {
	event, err = s.eventStorage.GetByID(id)
	return
}
