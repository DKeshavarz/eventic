package usecase

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	usecase "github.com/DKeshavarz/eventic/internal/usecase/guest"
	"github.com/stretchr/testify/assert"
)

func TestLoginWithEmail(t *testing.T) {
	testCases := []struct {
		name      string
		email     string
		password  string
		setupMock func(m *userStorage)
		wantErr   error
		wantUser  *entity.User
	}{
		{
			name: "valid email and password",
			email: "danny@gmail.com",
			password: "1234",
			setupMock: func(m *userStorage) {
				m.On("GetUserByEmail", "danny@gmail.com").Return(&entity.User{
					ID:       1,
					Email:    strPtr("danny@gmail.com"),
					Password: "1234",
				}, nil)
			},
			wantErr:  nil,
			wantUser: &entity.User{
				ID:       1,
				Email:    strPtr("danny@gmail.com"),
				Password: "1234",
			},
		},
		{
			name: "invalid email",
			email: "dann.y@gmail.com",
			password: "1234",
			setupMock: func(m *userStorage) {
				m.On("GetUserByEmail", "danny").Return(nil, nil)
			},
			wantErr:  usecase.ErrInvalidEmail,
			wantUser: nil,
		},
		{
			name: "invalid password",
			email: "danny@gmail.com",
			password: "1234",
			setupMock: func(m *userStorage) {
				m.On("GetUserByEmail", "danny@gmail.com").Return(&entity.User{
					ID:       1,
					Email:    strPtr("danny@gmail.com"),
					Password: "1111",
				}, nil)
			},
			wantErr:  usecase.ErrInvalidPassword,
			wantUser: nil,
		},
		{
			name: "not found user",
			email: "danny@gmail.com",
			password: "1234",
			setupMock: func(m *userStorage) {
				m.On("GetUserByEmail", "danny@gmail.com").Return(&entity.User{},
					usecase.ErrUserNotFound,
				)
			},
			wantErr:  usecase.ErrUserNotFound,
			wantUser: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			userStorage := new(userStorage)

			tc.setupMock(userStorage)

			guest := usecase.NewGuest(userStorage)

			user, err := guest.LoginWtihEmail(tc.email, tc.password)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantUser, user)
		})
	}
}
