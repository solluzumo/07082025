package model

import "time"

type TaskStatus string

const (
	StatusWaiting     TaskStatus = "waiting"
	StatusDownloading TaskStatus = "downloading files"
	StatusArchive     TaskStatus = "archive"
	StatusFailed      TaskStatus = "failed"
)

type TaskObject struct {
	ID        string
	Status    TaskStatus
	CreatedAt string
	Error     string
}

func NewTaskObject(id string) *TaskObject {
	return &TaskObject{
		ID:        id,
		Status:    StatusWaiting,
		CreatedAt: time.Now().Format(time.RFC3339),
		Error:     "",
	}
}
