package test

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/stretchr/testify/mock"
)

type mockUserStorage struct {
	mock.Mock
}

// ------- helpers ----------------
func strPtr(s string) *string {
	return &s
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
