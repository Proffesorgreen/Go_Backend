package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
	"github.com/zaahidali/task_manager_api/middleware"
)

func Router_Task() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.HandleRegister)
	router.POST("/login", controllers.HandleLogin)

	auth := router.Group("/tasks")
	auth.Use(middleware.AuthMiddleWare())

	auth.GET("/", controllers.HandleGetAllTasks)
	auth.GET("/:id", controllers.HandleGetTaskById)
	auth.PUT("/:id", middleware.RoleMiddleware("admin"), controllers.HandleUpdateTask)
	auth.DELETE("/:id", middleware.RoleMiddleware("admin"), controllers.HandleDeleteTask)
	auth.POST("/", middleware.RoleMiddleware("admin"), controllers.HandleCreateTask)

	return router
}
