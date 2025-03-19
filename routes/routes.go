package routes

import (
	"go_todolist/handlers"
	"go_todolist/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	api := gin.Default()
	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)
	// api.POST("/logout", handlers.Logout)
	// api.POST("/add", handlers.Add)
	// api.POST("/delete", handlers.Delete)
	// api.POST("/update", handlers.Update)
	// api.POST("/get", handlers.Get)
	authorized := api.Group("/")
	authorized.Use(middleware.AuthMiddleWare())
	{
		authorized.GET("/tasks", handlers.GetTasks)
	}
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return api
}
