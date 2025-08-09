package service

import (
	"07082025/internal/domain/model"
	"07082025/internal/domain/repository"
	"context"
	"time"
)

type TaskService struct {
	sem      repository.SemaphoreRepository
	taskRepo repository.TaskRepository
}

func NewTaskService(semRepo repository.SemaphoreRepository, tRepo repository.TaskRepository) *TaskService {
	return &TaskService{
		sem:      semRepo,
		taskRepo: tRepo,
	}
}

func (ts *TaskService) StartTask(ctx context.Context, taskID string) error {
	t := model.NewTaskObject(taskID)
	if err := ts.taskRepo.Store(t); err != nil {
		return err
	}
	if err := ts.sem.Acquire(ctx); err != nil {
		t.Status = model.StatusFailed
		t.Error = "failed to acquire"

		return err
	}
	defer ts.sem.Release()

	go ts.ExecuteTask(ctx, t)
	return nil
}

func (ts *TaskService) ExecuteTask(ctx context.Context, t *model.TaskObject) {
	defer ts.sem.Release()

	t.Status = model.StatusDownloading
	ts.taskRepo.Store(t)

	for i := 0; i <= 100; i += 10 {
		select {
		case <-ctx.Done():
			t.Status = model.StatusFailed
			t.Error = ctx.Err().Error()
			ts.taskRepo.Store(t)
			return
		default:
			ts.taskRepo.Store(t)
			time.Sleep(1 * time.Second)
		}
	}

	t.Status = model.StatusArchive
	ts.taskRepo.Store(t)
}

func (ts *TaskService) GetTaskStatus(ctx context.Context, id string) (*model.TaskObject, error) {
	return ts.taskRepo.FindById(id)
}
