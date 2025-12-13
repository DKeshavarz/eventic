package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type TokenSigner interface {
	Sign(claims jwt.Claims) (string, error)
	Parse(tokenString string, claims jwt.Claims) (*jwt.Token, error)
}

const (
	TokenTypeAccess       = "access"
	TokenTypeConfirmation = "confirmation"
)

// ------------------------- imp ------------------------

func NewHMACSigner(secret []byte) TokenSigner {
	return &hmacSigner{secret: secret}
}

type hmacSigner struct {
	secret []byte
}

func (s *hmacSigner) Sign(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *hmacSigner) Parse(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	token , err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})
	return token, err
}
