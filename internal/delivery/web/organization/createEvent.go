package organization

import (
	"net/http"
	"strconv"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/pkg/utile"
	"github.com/gin-gonic/gin"
)

type CreateEventRequest struct {
	Title       string     `json:"title"`
	Cost        int        `json:"cost"`
	DateTime    *time.Time `json:"datetime"`
	Description string     `json:"description"`
	Location    *string    `json:"location"`
	PosterPic   *string    `json:"poster_pic"`
	Link        *string    `json:"link"`
}

type CreateEventResponse struct {
	CreatedEvent *entity.Event `json:"created_event"`
}

// @Description A company/organization creates an event, providing title, date, location, etc.
// @Tags Organization
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Param CreateEventRequest body CreateEventRequest true "Event creation payload"
// @Security    BearerAuth
// @Success 201 {object} CreateEventResponse
// @Failure 400 {object} ErrorResponse "Invalid request or invalid organization ID"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /organization/{id}/event [post]
func (h *Handler) CreateEvent(c *gin.Context) {
	var req CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, DefaultErr(err.Error()))
		return
	}

	orgIDStr := c.Param("id")
	orgID, err := strconv.Atoi(orgIDStr)
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

	// Create by uscases
	if req.PosterPic != nil {
		if exsit := utile.FileExists(*req.PosterPic); !exsit {
			c.JSON(http.StatusBadRequest, DefaultErr("files doesnt exist"))
			return
		}
	}
	event := &entity.Event{
		OrganizerID: orgID,
		Title:       req.Title,
		Cost:        req.Cost,
		DateTime:    req.DateTime,
		Description: req.Description,
		Location:    req.Location,
		PosterPic:   req.PosterPic,
		Link:        req.Link,
	}

	newEvent, err := h.eventService.Create(userID, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultErr(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, CreateEventResponse{
		CreatedEvent: newEvent,
	})
}
