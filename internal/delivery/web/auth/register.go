package auth

import (
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/usecase/auth"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService         user.Service
	AuthService         auth.Service
	TokenSevice         jwt.AccessTokenService
	RefreshTokenService jwt.AccessTokenService
	SignupToken         jwt.SignupToken
}

func NewHandler(UserService user.Service, TokenSevice jwt.AccessTokenService, RefreshTokenService jwt.AccessTokenService, AuthService auth.Service, signupToken jwt.SignupToken) *Handler {
	return &Handler{
		UserService:         UserService,
		TokenSevice:         TokenSevice,
		RefreshTokenService: RefreshTokenService,
		AuthService:         AuthService,
		SignupToken:         signupToken,
	}
}

func RegisterRoutes(group *gin.RouterGroup, h *Handler) {
	group.POST("/login-email", h.LoginWithEmail)
	group.POST("/login-phone", h.LoginWithPhone)
	group.POST("/refresh-token", h.RefreshToken)
	group.POST("/signup-otp-request", h.SendEmailOTP)
	group.POST("/signup-otp-verify", h.VerifyEmailOTP)
}
