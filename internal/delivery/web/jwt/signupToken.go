package jwt

import (
	"errors"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity/validation"
	"github.com/golang-jwt/jwt/v5"
)

type SignupToken interface {
	Generate(eamil string) (string, error)
	Validate(tokenString string) (*SignupTokenClaims, error)
}

type SignupTokenConfig struct {
	Duration time.Duration
	Secret   []byte
}

type SignupTokenClaims struct {
	Email string `json:"Email"`
	jwt.RegisteredClaims
}

// ------------------------- imp -------------------------

type signupTokenService struct {
	TokenSigner
	duration time.Duration
}

func NewSignupTokenService(cfg *SignupTokenConfig) SignupToken {
	return &signupTokenService{
		TokenSigner: NewHMACSigner(cfg.Secret),
		duration:    cfg.Duration,
	}
}

func (s *signupTokenService) Generate(email string) (string, error) {
	if err := validation.Email(email); err != nil {
		return "", err
	}

	now := time.Now()
	claims := &SignupTokenClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.duration)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	return s.TokenSigner.Sign(claims)
}
func (s *signupTokenService) Validate(tokenString string) (*SignupTokenClaims, error) {
	claims := &SignupTokenClaims{}

	token, err := s.TokenSigner.Parse(tokenString, claims)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
