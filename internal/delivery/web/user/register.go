package user

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func NewHandler() *UserHandler {
	return &UserHandler{}
}

func RegisterRoutes(group *gin.RouterGroup, userHandler *UserHandler) {

}
