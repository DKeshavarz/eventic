package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginWithPhoneRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

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
	Meta  string `json:"meta,omitempty"`
}

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

// LoginWithPhone godoc
// @Summary     Login a user with Phone
// @Description Login a user with Phone to app and generate a JWT token and a refresh token
// @Tags        Auth
// @Accept      json
// @Param       login body LoginWithPhoneRequest true "Login info"
// @Produce     json
// @Success     200 {object} LoginResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /login-phone [POST]
func (h *Handler) LoginWithPhone(c *gin.Context) {
	// Get login request
	var req LoginWithPhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	// Login with user
	user, err := h.UserService.LoginWithPhone(req.Phone, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// Generate token for this user
	token, err := h.TokenSevice.Generate(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token"})
		return
	}
	refreshToken, err := h.RefreshTokenService.Generate(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate refresh token"})
		return
	}

	// Return the results
	c.JSON(http.StatusOK, LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

// RefreshToken godoc
// @Summary     Refresh JWT token
// @Description Refresh the JWT token using a valid refresh token
// @Tags        Auth
// @Accept      json
// @Param       refresh body RefreshTokenRequest true "Refresh token"
// @Produce     json
// @Success     200 {object} LoginResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /refresh-token [POST]
func (h *Handler) RefreshToken(c *gin.Context) {
	// Get refresh token request
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	// Validate and generate new token
	claims, err := h.RefreshTokenService.Validate(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid refresh token"})
		return
	}

	user, err := h.UserService.GetByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token", Meta: err.Error()})
		return
	}
	token, err := h.TokenSevice.Generate(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token", Meta: err.Error()})
		return
	}

	// Return the new token
	c.JSON(http.StatusOK, LoginResponse{
		Token:        token,
		RefreshToken: req.RefreshToken, // Reuse the same refresh token
	})
}
