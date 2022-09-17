package controllers

import (
	"net/http"

	"github.com/bigpaulie/go-todo/pkg/controllers/requests"
	"github.com/bigpaulie/go-todo/pkg/models"
	"github.com/gin-gonic/gin"
)

func IndexTodo(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func ShowTodo(c *gin.Context) {
	var id string = c.Param("id")

	todo, err := models.FindTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var request = requests.CreateTodoRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	todo, err := models.CreateTodo(request.Task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var id string = c.Param("id")
	var request = requests.UpdateTodoRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	todo, err := models.UpdateTodo(id, request)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	var id string = c.Param("id")
	err := models.DeleteTodo(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Item deleted successfully.",
	})
}
