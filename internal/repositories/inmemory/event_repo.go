package inmemory

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type EventStorage struct {
	db *DB
}

func NewEventStorage(db *DB) repositories.Event {
	return &EventStorage{
		db: db,
	}
}

func (s *EventStorage) GetByID(id int) (*entity.Event, error) {
	s.db.mu.RLock()
	defer s.db.mu.RUnlock()

	if val, exist := s.db.events[id] ; exist {
		return val, nil
	}
	return nil, repositories.ErrEventNotFound
}
func (s *EventStorage) Create(event *entity.Event) (*entity.Event, error) {
	s.db.mu.Lock()
	defer s.db.mu.Unlock()

	event.ID = s.db.eventCounter
	s.db.events[event.ID] = event
	s.db.eventCounter++

	return event, nil
}
func (s *EventStorage)GetAll() ([]*entity.Event, error){
	return nil, nil
}