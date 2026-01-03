package organization

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateEventResquest struct {
}

type CreateEventResponse struct {
}

// @Router /organization/{id}/event [post]
func (h *Handler) CreateEvent(c *gin.Context) {
	var req CreateEventResquest
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

	userID++
	// h.eventService.Create(userID,)
	c.JSON(http.StatusCreated, orgID)
}
