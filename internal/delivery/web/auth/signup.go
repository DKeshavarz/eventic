package auth

import (
	"log"
	"net/http"

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
	log.Println("\n\n\n", claims.Email, "\n\n\n", req.Email)
	if req.Email == nil || claims.Email != *req.Email {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "مشکلی پیش آمده",
			Meta:  "token doesn't match the credential"},
		)
		return
	}

	

	c.JSON(http.StatusOK, SignUpResponse{
		Token: "mt.token",
		RefreshToken: "te.token",
	})
}
