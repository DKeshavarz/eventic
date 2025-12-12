package jwt_test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	token, err := jwtService.Generate(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateJWT(t *testing.T) {
	// Generate a token
	token, err := jwtService.Generate(user)
	assert.NoError(t, err)

	// Validate the token
	claims, err := jwtService.Validate(token)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, claims.UserID)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"
	claims, err := jwtService.Validate(invalidToken)
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateJWT_ExpiredToken(t *testing.T) {

	token, err := InvalidService.Generate(user)
	assert.NoError(t, err)

	claims, err := InvalidService.Validate(token)
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestInvalidSignature(t *testing.T) {
	jwtService2 := jwt.NewTokenService(&jwt.Config{
		Duration: time.Hour,
		Secret:   []byte("Meowwww"),
	})

	token, err := jwtService.Generate(user)
	assert.NoError(t, err)

	claims, err := jwtService2.Validate(token)
	assert.Error(t, err)
	assert.Nil(t, claims)

	claims, err = jwtService.Validate(token)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
}
