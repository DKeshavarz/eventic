package auth

import (
	"net/http"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Token    string  `json:"token"`
}

type SignUpResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "مشکلی پیش آمده",
			Meta:  err.Error(),
		})
		return
	}

	claims, err := h.SignupToken.Validate(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "مشکلی پیش آمده", Meta: err.Error()})
		return
	}

	if req.Email == nil || claims.Email != *req.Email {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "مشکلی پیش آمده",
			Meta:  "token doesn't match the credential"},
		)
		return
	}

	user := &entity.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	newUser, err := h.UserService.Signup(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: err.Error()})
		return
	}

	token, err := h.TokenSevice.Generate(newUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "مشکلی پیش آمده",
			Meta: err.Error(),
		})
		return
	}

	refreshToken, err := h.RefreshTokenService.Generate(newUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "مشکلی پیش آمده",
			Meta: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SignUpResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}
