package event

import (
	"github.com/DKeshavarz/eventic/internal/entity"
)

func (s *service) Create(userID int,event *entity.Event) (*entity.Event, error) {
	if err := event.Validate(); err != nil {
		return nil, err
	}

	eventOrg, err := s.organizationStorage.GetByID(event.OrganizerID)
	if err != nil {
		return nil, err
	}
	
	if eventOrg.OwnerID != userID {
		return nil, ErrInvalidEventCreator
	}

	newEvent, err := s.eventStorage.Create(event)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}

func (s *service) Join(joinEvent *entity.JoinEvent) (*entity.JoinEvent, error) {
	if _, err := s.userStorage.GetByID(joinEvent.UserID); err != nil {
		return nil, err
	}

	if _, err := s.eventStorage.GetByID(joinEvent.EventID); err != nil {
		return nil, err
	}
	
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
