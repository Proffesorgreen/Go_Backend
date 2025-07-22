package main

import (
	"github.com/zaahidali/task_manager_api/routes"
)

func main() {
	router := routes.Router_Task()
	router.Run()
}
