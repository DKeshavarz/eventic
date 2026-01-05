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
// getOrganisations returns the list of organisations the given user is part of
// @Summary Get user's organisations
// @Description Retrieves all organisations/companies that the specified user is a member of.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} getOrganisationResponse "List of organisations the user belongs to"
// @Failure 400 {object} ErrorResponse "Invalid user ID"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /users/{id}/organisations [get]
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
