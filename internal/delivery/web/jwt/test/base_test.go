package jwt_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomHMACSigner_SignAndParse(t *testing.T) {
	type CustomClaims struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
		jwt.RegisteredClaims
	}

	claims := &CustomClaims{
		UserID: "user-123",
		Role:   "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user-123",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour)},
		},
	}

	tokenString, err := signerHMAC.Sign(claims)
	require.NoError(t, err)

	parsedClaims := &CustomClaims{}
	_, err = signerHMAC.Parse(tokenString, parsedClaims)
	require.NoError(t, err)

	assert.Equal(t, "user-123", parsedClaims.UserID)
	assert.Equal(t, "admin", parsedClaims.Role)
	assert.Equal(t, "user-123", parsedClaims.Subject)
}
