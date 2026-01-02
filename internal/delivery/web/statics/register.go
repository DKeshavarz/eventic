package statics

import (
	"os"

	"github.com/gin-gonic/gin"
)

const (

)

func Register(group *gin.RouterGroup) {
	os.MkdirAll("./static/upload", 0755)

	group.Static("/upload", "./static/upload")
	group.POST("/upload", uploadImage)
}