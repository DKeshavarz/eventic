package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidEmail(t *testing.T) {
	testCases := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "test@gmail.com",
			wantErr: false,
		},
		{
			name:    "invalid email",
			email:   "test@gmail",
			wantErr: true,
		},
		{
			name:    "test with some domain",
			email:   "test9090@something.ir",
			wantErr: false,
		},
		{
			name:    "email with extra",
			email:   "test@g.mail.com",
			wantErr: true,
		},
		{
			name:    "email with invalid domain",
			email:   "test@g.mail.c",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateEmail(tc.email)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_ValidPhone(t *testing.T) {
	testCases := []struct {
		name    string
		phone   string
		wantErr bool
	}{
		{
			name:    "valid phone",
			phone:   "09024066589",
			wantErr: false,
		},
		{
			name:    "less than 11 digits",
			phone:   "0912345678",
			wantErr: true,
		},
		{
			name:    "more than 11 digits",
			phone:   "091234567890",
			wantErr: true,
		},
		{
			name:    "contain not numbers",
			phone:   "0912345678a",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePhone(tc.phone)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
