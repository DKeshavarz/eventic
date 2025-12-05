package inmemory

import (
	"fmt"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type JoinEventStorage struct {
	db *DB
}

func NewJoinEventStorage(db *DB) repositories.JoinEvent {
	return &JoinEventStorage{
		db: db,
	}
}

func (s *JoinEventStorage) GetByUserID(id int) ([]*entity.JoinEvent, error) {
	s.db.mu.RLock()
	defer s.db.mu.RUnlock()
	list := make([]*entity.JoinEvent, 0)
	for _, value := range s.db.joinEvents {
		if value.UserID == id {
			list = append(list, value)
		}
	}
	return list, nil
}
func (s *JoinEventStorage) Create(joinEvent *entity.JoinEvent) (*entity.JoinEvent, error) {
	s.db.mu.Lock()
	defer s.db.mu.Unlock()
	key := fmt.Sprintf("%d-%d", joinEvent.EventID, joinEvent.UserID)
	s.db.joinEvents[key] = joinEvent
	return joinEvent, nil
}
