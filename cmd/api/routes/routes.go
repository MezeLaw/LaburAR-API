package routes

import (
	"LaburAR/cmd/api/auth"
	"LaburAR/cmd/api/middleware"
	"LaburAR/cmd/api/profile"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	authentication := router.Group("/v1/auth")
	{
		authentication.POST("/login", auth.Login)
		authentication.POST("/signup", auth.Signup)
	}

	prof := router.Group("/v1/settings")
	{
		prof.GET("/profile", middleware.AuthValidation(), profile.Profile)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
}
