package web

import (
	"net/http"
	"time"

	_ "github.com/DKeshavarz/eventic/docs"
	"github.com/DKeshavarz/eventic/internal/delivery/web/auth"
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title                      Eventic
// @version                    0.1.0
// @description                A platform to hold and participate in events
// @termsOfService             http://swagger.io/terms/
// @contact.name               Eventic Dev Team
// @contact.url                https://github.com/DKeshavarz/eventic
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type `Bearer ` followed by your JWT token. example: "Bearer abcde12345"
func Start(cfg *Config, userService user.Service) error {
	server := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DocExpansion("none"),
	)

	server.Use(cors.New(corsConfig))
	server.GET("/swagger/*any", swaggerHandler)
	server.GET("/health", health)

	token := jwt.NewSevice(cfg.Token)
	refreshToken := jwt.NewSevice(cfg.RefreshToken)

	authHandler := auth.NewHandler(userService, token, refreshToken)
	auth.RegisterRoutes(server.Group(""), authHandler)
	return server.Run(":" + cfg.Port)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"time":    time.Now(),
		"service": "eventic v0.0.0",
	})
}
