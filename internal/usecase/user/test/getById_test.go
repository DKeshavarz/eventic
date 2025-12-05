package usecase

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	UserWithID57 := &entity.User{
		ID:       57,
		Username: "user 57",
		Password: "password",
	}

	testCases := []struct {
		title     string
		id        int
		setupMock func(m *userStorage)
		wantErr   error
		wantUser  *entity.User
	}{
		{
			title: "Existing user",
			id:    57,
			setupMock: func(m *userStorage) {
				m.On("GetByID", 57).Return(UserWithID57, nil)
			},
			wantErr:  nil,
			wantUser: UserWithID57,
		},
		{
			title: "Non existing user",
			id:    65,
			setupMock: func(m *userStorage) {
				m.On("GetByID", 65).Return(&entity.User{}, repositories.ErrEventNotFound)
			},
			wantErr:  repositories.ErrEventNotFound,
			wantUser: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {

			userStorage := new(userStorage)

			tc.setupMock(userStorage)

			userSevice := user.NewSevice(userStorage)

			user, err := userSevice.GetByID(tc.id)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantUser, user)
		})
	}

}
