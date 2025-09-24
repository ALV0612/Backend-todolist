package routes

import (
	"todo-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// API group
	api := r.Group("/api")
	{
		api.GET("/todos", controllers.GetTodos)
		api.GET("/todos/:id", controllers.GetTodo)
		api.POST("/todos", controllers.CreateTodo)
		api.PUT("/todos/:id", controllers.UpdateTodo)
		api.DELETE("/todos/:id", controllers.DeleteTodo)
	}
}
