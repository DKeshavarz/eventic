package web

import (
	"net/http"
	"time"

	"github.com/DKeshavarz/eventic/internal/delivery/web/auth"
	"github.com/DKeshavarz/eventic/internal/delivery/web/jwt"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(cfg *Config, userService user.Service) error {
	server := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}

	server.Use(cors.New(corsConfig))
	server.GET("/health", health)

	token := jwt.NewSevice(&jwt.Config{
		Duration: time.Minute * 30,
		Secret: []byte("meowwww"),
	})
	refreshToken := jwt.NewSevice(&jwt.Config{
		Duration: time.Hour * 5,
		Secret: []byte("meowwwwww"),
	})
	
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
