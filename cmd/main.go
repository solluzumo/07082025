package main

import (
	"07082025/internal/app"
	"07082025/internal/config"
	taskInfrastructure "07082025/internal/infrastructure"
)

func main() {
	cfg := config.Load()

	taskSemaphore := taskInfrastructure.NewChannelTask(cfg.Task.MaxTasksAmount)

	app := app.NewApp(taskSemaphore)
	app.Run()
}
