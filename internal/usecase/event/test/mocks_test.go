package test

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/stretchr/testify/mock"
)

// ---------------------------------------------
type mockJoinEventStorage struct {
	mock.Mock
}


func (e *mockJoinEventStorage) GetByUserID(id int) ([]*entity.JoinEvent, error) {
	args := e.Called(id)
	return args.Get(0).([]*entity.JoinEvent), args.Error(1)
}

// ---------------------------------------------

type mockEventStorage struct {
	mock.Mock
}

func (e *mockEventStorage) Create(event *entity.Event) (*entity.Event, error) {
	args := e.Called(event)
	return args.Get(0).(*entity.Event), args.Error(1)
}

func (e *mockEventStorage) GetByID(id int) (*entity.Event, error) {
	args := e.Called(id)
	return args.Get(0).(*entity.Event), args.Error(1)
}
func (e *mockEventStorage) GetAll() ([]*entity.Event, error) {
	args := e.Called()
	return args.Get(0).([]*entity.Event), args.Error(1)
}

// ---------------------------------------------
type mockOrganizionStorage struct {
	mock.Mock
}

func (m *mockOrganizionStorage) Create(org *entity.Organization) (*entity.Organization, error) {
	args := m.Called(org)
	return args.Get(0).(*entity.Organization), args.Error(1)
}

func (m *mockOrganizionStorage) GetByID(id int) (*entity.Organization, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Organization), args.Error(1)
}

func (m *mockOrganizionStorage) GetByOwnerID(id int) ([]*entity.Organization, error) {
	args := m.Called(id)
	return args.Get(0).([]*entity.Organization), args.Error(1)
}