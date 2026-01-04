package event

import (
	"net/http"
	"strconv"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
)

type joinEventResponse struct {
    UserId  int `json:"user_id" example:"42"`   
    EventID int `json:"event_id" example:"101"`
}

// CreateEvent registers the currently authenticated user for a given event
// @Summary Register for an event
// @Description Adds the current authenticated user as a participant in the specified event.
// @Tags Event 
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Security    BearerAuth
// @Success 201 {object} joinEventResponse "Successfully registered"
// @Failure 400 {object} ErrorResponse "Invalid event ID"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /event/{id}/registrations [post]
func (h *Handler) JoinEvent(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "آیدی رویداد نامعتبر است",
			Meta:  err.Error(),
		})
	}
	// Get User ID
	strUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, DefaultErr("user Id it token is missed"))
		return
	}

	userID, ok := strUserID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, DefaultErr("Can't cast user userID to int"))
		return
	}

	joinEvent := &entity.JoinEvent{
		EventID: eventID,
		UserID:  userID,
	}

	newJoinEvent, err := h.eventSerivce.Join(joinEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultErr(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, joinEventResponse{
		UserId: newJoinEvent.UserID,
		EventID: newJoinEvent.EventID,
	})
}
