package httpInterface

import (
	"07082025/internal/application/dto"
	service "07082025/internal/application/services"
	taskService "07082025/internal/application/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type TaskHandler struct {
	taskService *taskService.TaskService
}

func NewHTTPHandler(taskService *taskService.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Methow not allowed", http.StatusMethodNotAllowed)
	}

	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := h.taskService.StartTask(r.Context(), taskID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *TaskHandler) GetTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Methow not allowed", http.StatusMethodNotAllowed)
	}

	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	t, err := h.taskService.GetTaskStatus(r.Context(), taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func (h *TaskHandler) AddLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request dto.LinkRequestDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	err := service.AddLink(request)
	err.StatusCode{}
	fmt.Println(request)
}
