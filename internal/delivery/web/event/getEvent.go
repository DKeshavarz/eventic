package event

import (
	"net/http"
	"strconv"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
)

type GetEventResponse struct {
	Event *entity.Event `json:"event"`
}

// GetEvent godoc
// @Summary Get a single event by ID
// @Description Retrieve a specific event using its unique identifier
// @Tags Event
// @Accept json
// @Produce json
// @Param id path int true "Event ID" minimum(1)
// @Success 200 {object} GetEventResponse "Event retrieved successfully"
// @Failure 400 {object} ErrorResponse "Invalid event ID or bad request"
// @Failure 404 {object} ErrorResponse "Event not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /event/{id} [get]
func (h *Handler) GetEvents(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "آیدی رویداد نامعتبر است",
			Meta: err.Error(),
		})
		return
	}

	event, err := h.eventSerivce.Get(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "مشکلی پیش امده است",
			Meta: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,GetEventResponse{
		Event: event,
	})
}
