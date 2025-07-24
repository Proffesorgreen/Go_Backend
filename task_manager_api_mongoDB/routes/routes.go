package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
)

func Router_Task() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", controllers.HandleGetAllTasks)
	router.GET("/tasks/:id", controllers.HandleGetTaskById)
	router.PUT("/tasks/:id", controllers.HandeleUpdateTask)
	router.DELETE("/tasks/:id", controllers.HandleDeleteTask)
	router.POST("/tasks", controllers.HandleCreateTask)

	return router
}
