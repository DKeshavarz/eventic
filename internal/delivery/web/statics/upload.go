package statics

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadImageResponse struct {
	URL string `json:"url" example:"http://localhost:8080/uploads/image123.png"`
}

// UploadImage godoc
// @Summary Upload an image
// @Description Uploads a picture and returns its public URL
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file"
// @Success 200 {object} UploadImageResponse
// @Failure 400 {object} map[string]string
// @Router /static/upload [post]
func uploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// generate unique filename: timestamp + original name
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	filePath := filepath.Join("uploads", filename)

	// Save file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UploadImageResponse{URL: "/static/" + filePath})
}
