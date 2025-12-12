package jwt_test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/entity"
)

var (
	jwtService     jwt.TokenService
	InvalidService jwt.TokenService
	user           *entity.User
	signerHMAC     jwt.TokenSigner
)

func TestMain(m *testing.M) {
	jwtService = jwt.NewTokenService(&jwt.Config{
		Duration: 1 * time.Hour,
		Secret:   []byte("meow-2025"),
	})

	InvalidService = jwt.NewTokenService(&jwt.Config{
		Duration: 1 * -time.Hour,
		Secret:   []byte("meow"),
	})

	user = &entity.User{ID: 123}

	signerHMAC = jwt.NewHMACSigner([]byte("very-secret-key"))
	m.Run()
}
