package test

import (
	"github.com/gin-gonic/gin"
	"github.com/valianx/videos/config"
	"github.com/valianx/videos/internal"
	"github.com/valianx/videos/internal/application/handlers"
)

func setupServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	config.ConnectDataBaseProduction()
	r.GET("/ping", pingEndpoint)

	r.POST("/login", internal.AuthMiddleware.LoginHandler)

	//sistema usuarios
	r.GET("/users", handlers.FindUsers)
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.FindUser)
	r.PATCH("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	return r
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
