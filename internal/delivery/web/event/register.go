package event

import (
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

func RegisterRoutes(group *gin.RouterGroup, h *Handler) {
	group.GET("/", h.GetAllEvents)
}
