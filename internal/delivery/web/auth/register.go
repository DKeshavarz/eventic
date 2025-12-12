package auth

import (
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService         user.Service
	TokenSevice         jwt.AccessTokenService
	RefreshTokenService jwt.AccessTokenService
}

func NewHandler(UserService user.Service, TokenSevice jwt.AccessTokenService, RefreshTokenService jwt.AccessTokenService) *Handler {
	return &Handler{
		UserService:         UserService,
		TokenSevice:         TokenSevice,
		RefreshTokenService: RefreshTokenService,
	}
}

func RegisterRoutes(group *gin.RouterGroup, h *Handler) {
	group.POST("/login-email", h.LoginWithEmail)
	group.POST("/login-phone", h.LoginWithPhone)
	group.POST("/refresh-token", h.RefreshToken)
}
