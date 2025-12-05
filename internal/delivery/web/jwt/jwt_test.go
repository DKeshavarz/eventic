package jwt

import (
    "testing"
    "time"

    "github.com/DKeshavarz/eventic/internal/entity"
    "github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
    jwtService := NewSevice(&Config{
        Duration: time.Hour,
        Secret: []byte("Meow"),
    })
    user := &entity.User{ID: 123}


    token, err := jwtService.Generate(user)
    assert.NoError(t, err)
    assert.NotEmpty(t, token)
}

func TestValidateJWT(t *testing.T) {
    jwtService := NewSevice(&Config{
        Duration: time.Hour,
        Secret: []byte("Meow"),
    })
    user := &entity.User{ID: 123}

    // Generate a token
    token, err := jwtService.Generate(user)
    assert.NoError(t, err)

    // Validate the token
    claims, err := jwtService.Validate(token)
    assert.NoError(t, err)
    assert.Equal(t, user.ID, claims.UserID)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
    jwtService := NewSevice(&Config{
        Duration: time.Hour,
        Secret: []byte("Meow"),
    })

    invalidToken := "invalid.token.string"
    claims, err := jwtService.Validate(invalidToken)
    assert.Error(t, err)
    assert.Nil(t, claims)
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
    jwtService := NewSevice(&Config{
        Duration: -time.Hour,
        Secret: []byte("Meow"),
    })
    user := &entity.User{ID: 123}

    token, err := jwtService.Generate(user)
    assert.NoError(t, err)

    claims, err := jwtService.Validate(token)
    assert.Error(t, err)
    assert.Nil(t, claims)
}

func TestInvalidSignature(t *testing.T) {
    jwtService1 := NewSevice(&Config{
        Duration: time.Hour,
        Secret: []byte("Meow"),
    })
    jwtService2 := NewSevice(&Config{
        Duration: time.Hour,
        Secret: []byte("Meowwww"),
    })
    user := &entity.User{ID: 123}

    token, err := jwtService1.Generate(user)
    assert.NoError(t, err)


    claims, err := jwtService2.Validate(token)
    assert.Error(t, err)
    assert.Nil(t, claims)

    claims, err = jwtService1.Validate(token)
    assert.NoError(t, err)
    assert.NotNil(t, claims)
}