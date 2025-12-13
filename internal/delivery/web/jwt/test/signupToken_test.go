package jwt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSignupToken(t *testing.T) {
	_, err := signupTokenService.Generate("badMail@mail")
	assert.NotNil(t, err)

	_, err = signupTokenService.Generate("GoodMail@mail.com")
	assert.Nil(t, err)
}
func TestBadValidate(t *testing.T) {
	token := "someRandomString"
	_, err := signupTokenService.Validate(token)
	assert.NotNil(t, err)
}
func TestGenerateAndValidate(t *testing.T) {
	token, err := signupTokenService.Generate("GoodMail@mail.com")
	assert.Nil(t, err)

	claim, err := signupTokenService.Validate(token)
	assert.Nil(t, err) 
	if assert.NotNil(t, claim) {
		assert.Equal(t, claim.Email, "GoodMail@mail.com")
	}
}
