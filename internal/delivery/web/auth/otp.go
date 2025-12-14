package auth

import (
	"net/http"
	"time"

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

	//TODO: email
	if err := h.AuthService.SendOTP(req.Email, 5*time.Second); err != nil {
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
