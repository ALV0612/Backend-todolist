package main

import (
	"log"
	"os"
	"todo-backend/database"
	"todo-backend/middleware"
	"todo-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()
	//set up gin
	r := gin.Default()
	//set up cors
	r.Use(middleware.SetupCORS())
	//set up routes
	routes.SetupRoutes(r)
	//health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "🚀 Todo API is running!",
			"version":  "1.0.0",
			"database": "PostgreSQL",
			"status":   "healthy",
		})
	})
	//run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	log.Printf("🚀 Server starting on http://localhost:%s", port)
	log.Printf("📝 API endpoints: http://localhost:%s/api/todos", port)
	r.Run(":" + port)
}
