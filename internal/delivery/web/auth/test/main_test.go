package auth

import (
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	userSvc    *MockUserService
	tokenSvc   *MockJWTService
	refreshSvc *MockJWTService
	server     *gin.Engine
	group      *gin.RouterGroup
)

func TestMain(m *testing.M) {
	userSvc = new(MockUserService)
	tokenSvc = new(MockJWTService)
	refreshSvc = new(MockJWTService)

	server = gin.New()
	group = server.Group("/auth")

	m.Run()

}
