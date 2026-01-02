package organization

import (
	"net/http"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/gin-gonic/gin"
)

type CreateOrgRequest struct {
	org *entity.Organization
}

type CreateOrgResponse struct {
	org *entity.Organization
}

// CreateOrganization godoc
// @Summary Create new organization
// @Description Creates a new organization for the authenticated user
// @Tags organization
// @Accept json
// @Produce json
// @Param organization body CreateOrgRequest true "Organization data"
// @Success 201 {object} CreateOrgResponse
// @Failure 400 {object} ErrorResponse"Bad request"
// @Failure 500 {object} ErrorResponse"Internal server error"
// @Router /organization/ [post]
func (h *Handler) Create(c *gin.Context) {
	// Get user request
	var req CreateOrgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, DefaultErr(err.Error()))
		return
	}

	// Get User ID
	strUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, DefaultErr("user Id it token is missed"))
	}

	userID, ok := strUserID.(int) 
	if !ok {
		c.JSON(http.StatusInternalServerError, DefaultErr("Can't cast user userID to int"))
	}
	
	
	// Create by uscases
	req.org.OwnerID = userID
	org ,err := h.orgSevice.Create(req.org)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, DefaultErr(err.Error()))
		return
	}

	// Return resault
	c.JSON(http.StatusCreated, CreateOrgResponse{
		org: org,
	})
}
