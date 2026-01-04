package event

import (
	"net/http"
	"strconv"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
)

type joinEventResponse struct {
	UserId  int `json:"user_id"`
	EventID int `json:"event_id"`
}

// @Router /events/{id}/registrations [post]
func (h *Handler) CreateEvent(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "آیدی سازمان نامعتبر است",
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
