package router

import (
	"task_manager/delivery/controller"
	"task_manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func AllRouter(r *gin.Engine, allController *controller.AllController, auth *infrastructure.AuthMiddleware) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", allController.HandleRegister)
		userGroup.POST("/login", allController.HandleLogin)
	}

	taskGroup := r.Group("/tasks")
	taskGroup.Use(auth.Authentcate())
	{
		taskGroup.POST("/", auth.AuthorizeUser(), allController.HandleCreateTask)
		taskGroup.PATCH("/:id", auth.AuthorizeUser(), allController.HandleUpdateTask)
		taskGroup.DELETE("/:id", auth.AuthorizeUser(), allController.HandleDeleteTask)
		taskGroup.GET("/", allController.HandleGetAllTasks)
		taskGroup.GET("/:id", allController.HandleGetTaskById)
	}
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}
