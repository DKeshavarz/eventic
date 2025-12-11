package auth

import (
	"crypto/rand"
	"math/big"
)

func generateCode(length int) (string, error) {
	if length <= 0 {
		return "", nil 
	}

	result := make([]byte, length)
	for i := range length {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		result[i] = '0' + byte(n.Int64())
	}
	return string(result), nil
}