package routing

import (
	"github.com/bigpaulie/go-todo/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetupApiRouting(e *gin.Engine) {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/todo", controllers.IndexTodo)
		v1.POST("/todo", controllers.CreateTodo)
		v1.GET("/todo/:id", controllers.ShowTodo)
		v1.PUT("/todo/:id", controllers.UpdateTodo)
		v1.DELETE("/todo/:id", controllers.DeleteTodo)
	}
}
