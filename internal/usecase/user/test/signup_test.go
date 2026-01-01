package usecase

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/DKeshavarz/eventic/pkg/utile"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	tests := []struct {
		tag       string
		user      *entity.User
		setupMock func(m *mockUserStorage)
		wantErr   error
		wantUser  *entity.User
	}{
		{
			tag:       "Bad password",
			setupMock: func(m *mockUserStorage) {},
			user: &entity.User{
				Username: "muUser",
				Password: "pass",
			},
			wantErr: entity.ErrWeakPassword,
		},
		{
			tag:       "Valid",
			user: &entity.User{Username: "Gooduser", Password: "pass12345", Email: utile.StrPtr("dan@gmail.com") },
			setupMock: func(m *mockUserStorage) {
				m.On("Create", &entity.User{Username: "Gooduser", Password: "pass12345", Email: utile.StrPtr("dan@gmail.com")}).Return(
					&entity.User{ID: 1,Username: "Gooduser", Password: "pass12345", Email: utile.StrPtr("dan@gmail.com")},
					nil)
			},
			wantErr: nil,
			wantUser: &entity.User{ID: 1,Username: "Gooduser", Password: "pass12345", Email: utile.StrPtr("dan@gmail.com")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.tag, func(t *testing.T) {
			userStorage := new(mockUserStorage)
			tc.setupMock(userStorage)

			service := user.NewSevice(userStorage)
			user, err := service.Signup(tc.user)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantUser, user)
		})
	}
}
