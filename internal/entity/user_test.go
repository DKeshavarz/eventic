package entity

import (
	"testing"

	"github.com/DKeshavarz/eventic/pkg/utile"
)

func TestValidateUser(t *testing.T) {
	testCases := []struct {
		tag         string
		user        *User
		expectedErr error
	}{
		{
			tag: "week password",
			user: &User{
				Username: "ali",
				Password: "pass",
				Email: utile.StrPtr("ali@gmail.com"),
			},
			expectedErr: ErrWeakPassword,
		},
		{
			tag: "lack of credential",
			user: &User{
				Username: "ali",
				Password: "pass12345",
			},
			expectedErr: ErrNotEnoughCredential,
		},
		{
			tag: "Bad phone number",
			user: &User{
				Username: "ali",
				Password: "pass12345",
				Phone: utile.StrPtr("0938351437"),
			},
			expectedErr: ErrInvalidPhone,
		},
		{
			tag: "Bad email",
			user: &User{
				Username: "ali",
				Password: "pass12345",
				Email: utile.StrPtr("dan"),
			},
			expectedErr: ErrInvalidEmail,
		},
		{
			tag: "Valid",
			user: &User{
				Username: "ali",
				Password: "pass12345",
				Email: utile.StrPtr("dan@gmail.com"),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.tag, func(t *testing.T) {
			err := (tc.user).Validate()

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
