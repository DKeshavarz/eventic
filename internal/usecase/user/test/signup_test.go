package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	fmt.Println("meow")
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
			wantErr: errors.New("somthing"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.tag, func(t *testing.T) {
			userStorage := new(mockUserStorage)
			tc.setupMock(userStorage)

			service := user.NewSevice(userStorage)
			user, err := service.Signup(tc.user)
			assert.NotNil(t, nil)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantUser, user)
		})
	}
}
