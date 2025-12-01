package jwt

import (
    "testing"
    "time"

    "github.com/DKeshavarz/eventic/internal/entity"
    "github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
    user := &entity.User{ID: 123}
    duration := time.Hour

    token, err := Generate(user, duration)
    assert.NoError(t, err)
    assert.NotEmpty(t, token)
}

func TestValidateJWT(t *testing.T) {
    user := &entity.User{ID: 123}
    duration := time.Hour

    // Generate a token
    token, err := Generate(user, duration)
    assert.NoError(t, err)

    // Validate the token
    claims, err := Validate(token)
    assert.NoError(t, err)
    assert.Equal(t, user.ID, claims.UserID)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
    // Test with an invalid token
    invalidToken := "invalid.token.string"
    claims, err := Validate(invalidToken)
    assert.Error(t, err)
    assert.Nil(t, claims)
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
    user := &entity.User{ID: 123}
    duration := -time.Hour // Token already expired

    // Generate an expired token
    token, err := Generate(user, duration)
    assert.NoError(t, err)

    // Validate the expired token
    claims, err := Validate(token)
    assert.Error(t, err)
    assert.Nil(t, claims)
}