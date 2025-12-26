package auth

import (
	"net/http"

	"github.com/DKeshavarz/eventic/internal/entity/validation"
	"github.com/gin-gonic/gin"
)

type VerifyEmailOTPRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type VerifyEmailOTPResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}

// SendEmailOTP godoc
// @Summary     Validate otp for a user
// @Description Validate otp by email to a user for signup process of a user
// @Tags        Auth
// @Accept      json
// @Param       login body VerifyEmailOTPRequest true "Login info"
// @Produce     json
// @Success     200 {object} VerifyEmailOTPResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /signup-otp-verify [POST]
func (h *Handler) VerifyEmailOTP(c *gin.Context) {
	var req VerifyEmailOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	if err := validation.Email(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.AuthService.VerifyOTP(req.Email, req.Code); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	
	token, err := h.SignupToken.Generate(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Try again",
			Meta: err.Error(),
		})
	}

	c.JSON(http.StatusOK, VerifyEmailOTPResponse{
		Success: true,
		Token:   token,
		Message: "otp verified successfully",
	})
}
