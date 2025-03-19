package routes

import (
	"go_todolist/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	api := gin.Default()
	api.POST("/register", handlers.Register)
	// api.POST("/login", handlers.Login)
	// api.POST("/logout", handlers.Logout)
	// api.POST("/add", handlers.Add)
	// api.POST("/delete", handlers.Delete)
	// api.POST("/update", handlers.Update)
	// api.POST("/get", handlers.Get)
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return api
}
