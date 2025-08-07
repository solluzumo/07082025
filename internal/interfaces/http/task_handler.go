package httpInterface

import (
	taskService "07082025/internal/application/services"
	"07082025/internal/domain/model"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	taskService *taskService.TaskService
}

func NewHTTPHandler(taskService *taskService.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) StartTask(w http.ResponseWriter, r *http.Request) {
	taskID := model.TaskID(r.URL.Query().Get("id"))
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

func (h *TaskHandler) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	taskID := model.TaskID(r.URL.Query().Get("id"))
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
