package usecase

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userStorage struct {
	mock.Mock
}

func TestLoginWithPhone(t *testing.T) {
	testCases := []struct {
		name     string
		phone    string
		password string
		setupMock     func(m *userStorage)
		wantErr  error
		wantUser *entity.User
	}{
		{
			name:     "success 1",
			phone:    "09123456789",
			password: "123456",
			setupMock: func(m *userStorage) {
				m.On("GetUserByPhone", "09123456789").Return(&entity.User{
					ID:       1,
					Phone:    strPtr("09123456789"),
					Password: "123456",
				}, nil)
			},
			wantErr:  nil,
			wantUser: &entity.User{
				ID:       1,
				Phone:    strPtr("09123456789"),
				Password: "123456",
			},
		},
		{
			name:     "invalid number - long number",
			phone:    "091234567895",
			password: "123456",
			setupMock: func(m *userStorage) {
				m.On("GetUserByPhone").Return(nil, nil)
			},
			wantErr:  user.ErrInvalidPhone,
			wantUser: nil,
		},
		{
			name:     "invalid number - with invalid characters",
			phone:    "0912345678+",
			password: "123456",
			setupMock: func(m *userStorage) {
				m.On("GetUserByPhone").Return(nil, nil)
			},
			wantErr:  user.ErrInvalidPhone,
			wantUser: nil,
		},
		{
			name:     "success 2",
			phone:    "09188119090",
			password: "1111",
			setupMock: func(m *userStorage) {
				m.On("GetUserByPhone", "09188119090").Return(&entity.User{
					ID:       2,
					Phone:    strPtr("09188119090"),
					Password: "1111",
				}, nil)
			},
			wantErr:  nil,
			wantUser: &entity.User{
				ID:       2,
				Phone:    strPtr("09188119090"),
				Password: "1111",
			},
		},
		{
			name:     "user not found",
			phone:    "09188119091",
			setupMock: func(m *userStorage) {
				m.On("GetUserByPhone", "09188119091").Return(&entity.User{}, user.ErrUserNotFound)
			},
			password: "1111",
			wantErr:  user.ErrUserNotFound,
			wantUser: nil,
		},
		{
			name:     "invalid password",
			phone:    "09188119091",
			setupMock: func(m *userStorage) {
				m.On("GetUserByPhone", "09188119091").Return(&entity.User{
					ID:       2,
					Phone:    strPtr("09188119091"),
					Password: "2222",
				}, nil)
			},
			password: "1111",
			wantErr:  user.ErrInvalidPassword,
			wantUser: nil,
		},
	}


	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			userStorage := new(userStorage)

			tc.setupMock(userStorage)

			guest := user.NewGuest(userStorage)
			
			user, err := guest.LoginWithPhone(tc.phone, tc.password)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantUser, user)
		})
	}

}

// ------- helpers ----------------
func strPtr(s string) *string {
	return &s
}

func (u *userStorage) GetUserByPhone(phone string) (*entity.User, error) {
	args := u.Called(phone)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (u *userStorage) GetUserByEmail(email string) (*entity.User, error) {
	args := u.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (u *userStorage) Create(user *entity.User) (*entity.User, error) {
	args := u.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}
