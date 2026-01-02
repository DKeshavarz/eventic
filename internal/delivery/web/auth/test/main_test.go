package auth

import (
	"testing"
)

var (
	userSvc     *MockUserService
	tokenSvc    *MockJWTService
	refreshSvc  *MockJWTService
	signuptoken *MockSignupToken
)

func TestMain(m *testing.M) {
	userSvc = new(MockUserService)
	tokenSvc = new(MockJWTService)
	refreshSvc = new(MockJWTService)
	signuptoken = new(MockSignupToken)

	m.Run()

}
