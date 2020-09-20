package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hisami/go-todo-api/src/controller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/todo/api/v1")
	{
		v1.GET("/tasks", controller.TasksGET)
		v1.POST("/tasks", controller.TaskPOST)
	}
	router.Run(":9000")
}
