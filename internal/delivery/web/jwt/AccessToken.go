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

type AccessTokenClaims struct {
	UserID int `json:"userid"`
	jwt.RegisteredClaims
}


// ------------------------- imp -------------------------

type accessTokenService struct {
	TokenSigner
	duration time.Duration
}

func NewTokenService(cfg *AccessTokenConfig) AccessTokenService {
	return &accessTokenService{
		TokenSigner: NewHMACSigner(cfg.Secret),
		duration: cfg.Duration,
	}
}

func (s *accessTokenService) Generate(user *entity.User) (string, error) {
	now := time.Now()

	claims := &AccessTokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.duration)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	return s.TokenSigner.Sign(claims)
}

func (s *accessTokenService) Validate(tokenString string) (*AccessTokenClaims, error) {
	claims := &AccessTokenClaims{}

	token, err := s.TokenSigner.Parse(tokenString, claims)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	
	return claims, nil
}


