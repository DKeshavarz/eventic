package organization

import (
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/delivery/web/middelware"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/DKeshavarz/eventic/internal/usecase/organization"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	orgSevice organization.Service
	eventService event.Service
}

func NewHandler(orgSvc organization.Service, eventSvc event.Service) *Handler {
	return &Handler{
		orgSevice: orgSvc,
		eventService: eventSvc,
	}
}

func RegisterRoutes(group *gin.RouterGroup, h *Handler, accessToken jwt.AccessTokenService) {
	group.Use(middelware.Auth(accessToken))

	group.POST("/", h.Create)
	group.POST("/:id/event", h.CreateEvent)
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
