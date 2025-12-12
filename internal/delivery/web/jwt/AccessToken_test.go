package jwt

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/stretchr/testify/assert"
)

var (
	jwtService TokenService
    InvalidService TokenService
	user        *entity.User
)

func TestMain(m *testing.M) {
	jwtService = NewTokenService(&Config{
		Duration: 1 * time.Hour,
		Secret:   []byte("meow-2025"),
	})

    InvalidService = NewTokenService(&Config{
		Duration: 1 * -time.Hour,
		Secret:   []byte("meow"),
	})
	user = &entity.User{ID: 123}
	m.Run()
}
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
	jwtService2 := NewTokenService(&Config{
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
