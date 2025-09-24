package controllers

import (
	"net/http"
	"todo-backend/database"
	"todo-backend/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    todos,
		"count":   len(todos),
	})
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    todo,
	})
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Title is required",
		})
		return
	}

	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    todo,
	})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Todo not found",
		})
		return
	}

	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    todo,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Todo not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	database.DB.Save(&todo)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    todo,
	})
}
