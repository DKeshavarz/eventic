package usecase

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	m.Run()
}

// ---------------------------------------------------------
type mockUserStorage struct {
	mock.Mock
}

func (u *mockUserStorage) GetUserByPhone(phone string) (*entity.User, error) {
	args := u.Called(phone)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (u *mockUserStorage) GetUserByEmail(email string) (*entity.User, error) {
	args := u.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (u *mockUserStorage) Create(user *entity.User) (*entity.User, error) {
	args := u.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (u *mockUserStorage) GetByID(id int) (*entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

// ---------------------------------------------------------
type mockOrgStorage struct {
	mock.Mock
}

func (m *mockOrgStorage) Create(org *entity.Organization) (*entity.Organization, error) {
	args := m.Called(org)
	return args.Get(0).(*entity.Organization), args.Error(1)
}

func (m *mockOrgStorage) GetByID(id int) (*entity.Organization, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Organization), args.Error(1)
}

func (m *mockOrgStorage) GetByOwnerID(id int) ([]*entity.Organization, error) {
	args := m.Called(id)
	return args.Get(0).([]*entity.Organization), args.Error(1)
}
