package event

import (
	"net/http"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
)

type GetAllEventsResponse struct {
	events []*entity.Event
}
type ErrorResponse struct {
	Error string `json:"error"`
	Meta  string `json:"meta,omitempty"`
}

// Login        godoc
// @Summary     Get All Events
// @Description Get All Events
// @Tags        Event
// @Accept      json
// @Produce     json
// @Success     200 {object} GetAllEventsResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /event/ [get]
func (h *Handler) GetAllEvents(c *gin.Context) {
	events, err := h.eventSerivce.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Error: "مشکلی پیش آمده داداش",
				Meta:  err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, GetAllEventsResponse{
		events: events,
	})
}
