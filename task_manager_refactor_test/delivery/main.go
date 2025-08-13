package main

import (
	"task_manager/config"
	"task_manager/delivery/controller"
	"task_manager/delivery/router"
	"task_manager/infrastructure"
	"task_manager/repository"
	"task_manager/usecases"
)

func main() {
	config.LoadConfig()
	dbname := config.AppConfig.Dbname
	secret := config.AppConfig.SecretKey

	db := config.InitMongo().Database(dbname)
	usercollection := db.Collection("user_collection")
	taskcollection := db.Collection("Task_collection")

	userRepo := repository.NewUserRepoService(usercollection)
	taskRepo := repository.NewTaskRepoService(taskcollection)

	userPassword := infrastructure.NewPasswordProvider(10)
	userJWT := infrastructure.NewJwtProvider(secret)
	userAuth := infrastructure.NewAuthMiddleware(userJWT)

	userusecase := usecases.NewUserUseCase(userRepo, userPassword, userJWT)
	taskusecase := usecases.NewTaskUseCase(taskRepo)

	allcontroller := controller.NewAllController(userusecase, taskusecase)
	r := router.NewRouter()
	router.AllRouter(r, allcontroller, userAuth)

	if err := r.Run("localhost:8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
