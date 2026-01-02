package auth

import (
	"net/http"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity/validation"
	"github.com/gin-gonic/gin"
)

type SendEamilOTPRequest struct {
	Email string `json:"email"`
}

type method string

const (
	email method = "email"
)

type SendOTPResponse struct {
	Success         bool   `json:"success"`
	Message         string `json:"message"`
	Method          method `json:"method,omitempty"`            
	ExpiresInSecond int    `json:"expires_in_second,omitempty"` 

}

type VerifyEamilOTPRespone struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

// SendEmailOTP godoc
// @Summary     Send otp for a user
// @Description Send otp by email to a user for login process of a user
// @Tags        Auth
// @Accept      json
// @Param       login body SendEamilOTPRequest true "Login info"
// @Produce     json
// @Success     200 {object} SendOTPResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /signup-otp-request [POST]
func (h *Handler) SendEmailOTP(c *gin.Context) {
	var req SendEamilOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	//TODO: email duplicated
	if err := validation.Email(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.AuthService.SendOTP(req.Email, 2*time.Minute); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SendOTPResponse{
		Success: true,
		Message: "کد یکبار مصرف برای کاربر فرستاده شد",
		Method: email,
	})
}

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