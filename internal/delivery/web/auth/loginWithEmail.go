package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login        godoc
// @Summary     login a user with Eamil
// @Description Login a user with Eamil to app and generate a hwt token and a jwt refresh token
// @Tags        Auth
// @Accept      json
// @Param       login body LoginWithEmailRequest true "Login info"
// @Produce     json
// @Success     200 {object} LoginResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /login-email [POST]
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
		Token:        token,
		RefreshToken: refreshToken,
	})
}
