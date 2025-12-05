package inmemory

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type OrgStorage struct {
	db *DB
}

func NewOrgStorage(db *DB) repositories.Organization {
	return &OrgStorage{
		db: db,
	}
}

func (s *OrgStorage) GetByID(id int) (*entity.Organization, error) {
	s.db.mu.RLock()
	defer s.db.mu.RUnlock()

	if val, exist := s.db.organizations[id]; exist {
		return val, nil
	}
	return nil, repositories.ErrOrgNotFound
}

func (s *OrgStorage) Create(org *entity.Organization) (*entity.Organization, error) {
	s.db.mu.Lock()
	defer s.db.mu.Unlock()

	org.ID = s.db.orgCounter
	s.db.organizations[org.ID] = org
	s.db.orgCounter++

	return org, nil
}
