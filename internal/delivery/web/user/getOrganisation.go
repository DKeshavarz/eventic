package user

import (
	"net/http"
	"strconv"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
)

type getOrganisationResponse struct {
	UserID        int
	Organisations []*entity.Organization
}

// @Router /users/{id}/organisation [get]
func (h *Handler) getOrganisations(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "آیدی کاربر نامعتبر است",
			Meta:  err.Error(),
		})
		return
	}

	organisations, err := h.userService.GetCompanies(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultErr(err.Error()))
		return
	}

	c.JSON(http.StatusOK, getOrganisationResponse{
		Organisations: organisations,
		UserID: userID,
	})

}
