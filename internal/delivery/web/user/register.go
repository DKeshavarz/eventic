package user

import (
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService user.Service
}

func NewHandler(userSvc user.Service) *Handler {
	return &Handler{
		userService: userSvc,
	}
}

func RegisterRoutes(group *gin.RouterGroup, h *Handler){
	group.GET("/:id/organisations")
}

type ErrorResponse struct {
	Error string `json:"error"`
	Meta  string `json:"meta,omitempty"`
}

func DefaultErr(meta string) *ErrorResponse {
	return &ErrorResponse{
		Error: "مشکلی پیش آمده",
		Meta:  meta,
	}
}

