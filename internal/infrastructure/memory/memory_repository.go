package memory

import (
	"07082025/internal/domain/model"
	"07082025/internal/domain/repository"
	"errors"
	"sync"
)

var ErrTaskNotFound = errors.New("task with this id not found")

func NewMemoryRepository() repository.TaskRepository {
	return &MemoryRepository{
		tasks: make(map[string]*model.TaskObject),
	}
}

type MemoryRepository struct {
	mu    sync.RWMutex
	tasks map[string]*model.TaskObject
}

func (r *MemoryRepository) Store(task *model.TaskObject) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
	return nil
}

func (r *MemoryRepository) FindById(id string) (*model.TaskObject, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}
