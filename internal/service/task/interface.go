package task

import (
	"context"

	models "github.com/nottee-project/task_service/internal/models/task"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task models.CreateTaskParams) (models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) (models.UpdateTaskResponse, error)
	GetTask(ctx context.Context, task_id, user_id string) (models.GetTaskResponse, error)
	ListTasks(ctx context.Context, listTasksParams models.ListTasksParams) ([]models.GetTaskResponse, error)
	DeleteTask(ctx context.Context, task_id, user_id string) error
}
