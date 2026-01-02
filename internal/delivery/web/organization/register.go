package organization

import (
	"github.com/DKeshavarz/eventic/internal/usecase/organization"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	orgSevice organization.Service
}

func NewHandler(orgSvc organization.Service) *Handler {
	return &Handler{
		orgSevice: orgSvc,
	}
}

func RegisterRoutes(group *gin.RouterGroup, h *Handler) {
	group.POST("/", h.Create)
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
