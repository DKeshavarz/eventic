package organization

import (
	"net/http"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/pkg/utile"
	"github.com/gin-gonic/gin"
)

type CreateOrgRequest struct {
	Name        string  `json:"name" example:"TechMeet"`
	Description string  `json:"description" example:"Organization for tech events"`
	LogoPic     *string `json:"logo_pic,omitempty" example:"https://example.com/logo.png"`
	Email       *string `json:"email,omitempty" example:"info@company.com"`
	Phone       *string `json:"phone,omitempty" example:"+1234567890"`
}

type CreateOrgResponse struct {
	OrganizerID int     `json:"organizer_id" example:"1"`
	OwnerID     int     `json:"owner_id" example:"42"`
	Name        string  `json:"name" example:"TechMeet"`
	Description string  `json:"description" example:"Organization for tech events"`
	LogoPic     *string `json:"logo_pic,omitempty" example:"https://example.com/logo.png"`
	Email       *string `json:"email,omitempty" example:"info@company.com"`
	Phone       *string `json:"phone,omitempty" example:"+1234567890"`
}

// CreateOrganization godoc
// @Summary Create new organization
// @Description Creates a new organization for the authenticated user
// @Tags Organization
// @Accept json
// @Produce json
// @Param organization body CreateOrgRequest true "Organization data"
// @Security    BearerAuth
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
		return
	}

	userID, ok := strUserID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, DefaultErr("Can't cast user userID to int"))
		return
	}

	// Create by uscases
	if req.LogoPic != nil {
		if exsit := utile.FileExists(*req.LogoPic); !exsit {
			c.JSON(http.StatusBadRequest, DefaultErr("files doesnt exist"))
			return
		}
	}

	org, err := h.orgSevice.Create(&entity.Organization{
		OwnerID:     userID,
		Name:        req.Name,
		Description: req.Description,
		Email:       req.Email,
		Phone:       req.Phone,
		LogoPic:     req.LogoPic,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultErr(err.Error()))
		return
	}

	// Return resault
	c.JSON(http.StatusCreated, CreateOrgResponse{
		OrganizerID: org.ID,
		OwnerID:     org.OwnerID,
		Name:        org.Name,
		Description: org.Description,
		LogoPic:     org.LogoPic,
		Email:       org.Email,
		Phone:       org.Phone,
	})
}

// -------------------- Helpers -------------
