package event

import (
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/delivery/web/middelware"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	eventSerivce event.Service
}

func NewHandler(eventSerivce event.Service) *Handler {
	return &Handler{
		eventSerivce: eventSerivce,
	}
}

func RegisterRoutes(group *gin.RouterGroup, h *Handler, accessToken jwt.AccessTokenService) {
	group.GET("/", h.GetAllEvents)
	group.GET("/:id", h.GetEvents)

	group.POST("/:id/registrations", middelware.Auth(accessToken), h.JoinEvent)
}
