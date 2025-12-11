package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {
	testCases := []struct {
		tag   string
		lenth int
	}{
		{"lenth 6", 6},
		{"lenth 8", 8},
		{"lenth 10", 10},
	}

	for _, tc := range testCases {
		t.Run(tc.tag, func(t *testing.T) {
			otp, err := generateCode(tc.lenth)
			if len(otp) != tc.lenth {
				t.Errorf("Expected OTP length %d, got %d", tc.lenth, len(otp))
			}
			assert.Nil(t, err)
		})
	}
}
