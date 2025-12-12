package jwt

import (
	"errors"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenService interface {
	Generate(user *entity.User) (string, error)
	Validate(tokenString string) (*AccessTokenClaims, error)
}

type AccessTokenConfig struct {
	Duration time.Duration
	Secret   []byte
}
type tokenService struct {
	duration time.Duration
	secret   []byte
}

// ------------------------- imp -------------------------

type AccessTokenClaims struct {
	UserID int `json:"userid"`
	jwt.RegisteredClaims
}

func NewTokenService(cfg *AccessTokenConfig) AccessTokenService {
	return &tokenService{
		duration: cfg.Duration,
		secret:   cfg.Secret,
	}
}

func (s *tokenService) Generate(user *entity.User) (string, error) {
	claims := &AccessTokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(s.secret)
}

func (s *tokenService) Validate(tokenString string) (*AccessTokenClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
