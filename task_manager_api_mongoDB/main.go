package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/routes"
)

func main() {
	data.ConnectDB()
	setupGracefulShutdown()
	router := routes.Router_Task()
	router.Run()
}

func setupGracefulShutdown() {
	
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		data.DisconnectDB()
		os.Exit(0)
	}()
}
