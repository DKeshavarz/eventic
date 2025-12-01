package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginWithEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) LoginWithEmail(c *gin.Context) {
	// Get login reguest
	var req LoginWithEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	// login with user
	user, err := h.UserService.LoginWithEmail(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// generate token for this user
	token, err := h.TokenSevice.Generate(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token"})
		return
	}
	refreshToken, err := h.RefreshTokenService.Generate(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token"})
		return
	}

	// Return the resaults
	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
		RefreshToken: refreshToken,
	})
}
