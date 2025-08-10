package app

import (
	service "07082025/internal/application/services"
	"07082025/internal/domain/repository"
	"07082025/internal/infrastructure/memory"
	httpInterface "07082025/internal/interfaces/http"
	"net/http"
)

type App struct {
	taskRepo    repository.TaskRepository
	taskService *service.TaskService
	httpHandler *httpInterface.TaskHandler
}

func NewApp(semaphore repository.SemaphoreRepository) *App {
	taskRepo := memory.NewMemoryRepository()

	taskService := service.NewTaskService(semaphore, taskRepo)

	httpHandler := httpInterface.NewHTTPHandler(taskService)

	return &App{
		taskRepo:    taskRepo,
		taskService: taskService,
		httpHandler: httpHandler,
	}
}

func (a *App) Run() {
	http.HandleFunc("/task/start", a.httpHandler.StartTaskHandler)
	http.HandleFunc("/task/status", a.httpHandler.GetTaskStatusHandler)
	http.HandleFunc("/task/add-link", a.httpHandler.AddLinkHandler)
	http.ListenAndServe(":8080", nil)
}
