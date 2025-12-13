package jwt_test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/entity"
)

var (
	jwtService     jwt.AccessTokenService
	InvalidService jwt.AccessTokenService
	signupTokenService jwt.SignupToken
	user           *entity.User
	signerHMAC     jwt.TokenSigner
)

func TestMain(m *testing.M) {
	jwtService = jwt.NewTokenService(&jwt.AccessTokenConfig{
		Duration: 1 * time.Hour,
		Secret:   []byte("meow-2025"),
	})

	InvalidService = jwt.NewTokenService(&jwt.AccessTokenConfig{
		Duration: 1 * -time.Hour,
		Secret:   []byte("meow"),
	})

	signupTokenService = jwt.NewSignupTokenService(&jwt.SignupTokenConfig{
		Duration: 20 * time.Second,
		Secret: []byte("Meow-Meow-red"),
	})

	user = &entity.User{ID: 123}

	signerHMAC = jwt.NewHMACSigner([]byte("very-secret-key"))
	m.Run()
}
