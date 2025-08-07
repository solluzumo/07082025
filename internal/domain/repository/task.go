package repository

import "07082025/internal/domain/model"

type TaskRepository interface {
	Store(task *model.TaskObject) error
	FindById(taskID model.TaskID) (*model.TaskObject, error)
}
