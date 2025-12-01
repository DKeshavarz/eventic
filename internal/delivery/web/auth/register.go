package auth

import (
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService         user.Service
	TokenSevice         jwt.Service
	RefreshTokenService jwt.Service
}

func NewHandler() *Handler {
	return &Handler{}
}

func RegisterRoutes(group *gin.RouterGroup, h *Handler) {
	group.POST("/login-email", h.LoginWithEmail)
}
