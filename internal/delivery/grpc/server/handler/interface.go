package handler
import (
	"context"

	ssov1 "github.com/nottee-project/protos/fleap_protos"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *ssov1.CreateTaskRequest) (*ssov1.CreateTaskResponse, error)
	ListTasks(ctx context.Context, taskList *ssov1.ListTasksRequest) (*ssov1.ListTasksResponse, error)
	GetTask(ctx context.Context, task *ssov1.GetTaskRequest) (*ssov1.GetTaskResponse, error)
	UpdateTask(ctx context.Context, task *ssov1.UpdateTaskRequest) (*ssov1.UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, task *ssov1.DeleteTaskRequest) (*ssov1.DeleteTaskResponse, error)
}
