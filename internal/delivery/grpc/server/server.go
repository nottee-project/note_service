package grpc_server

import (
	"context"
	"log"

	ssov1 "github.com/nottee-project/protos/fleap_protos"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedTaskServiceServer
	taskService TaskService
	repo        TaskRepository
}

func Register(gRPCServer *grpc.Server, taskService TaskService) {
	ssov1.RegisterTaskServiceServer(gRPCServer, &serverAPI{taskService: taskService})
	log.Println("TaskService gRPC server registered")
}

type TaskService interface {
	CreateTask(ctx context.Context, task *ssov1.CreateTaskRequest) (*ssov1.CreateTaskResponse, error)
	ListTasks(ctx context.Context, tsaskList *ssov1.ListTasksRequest) (*ssov1.ListTasksResponse, error)
	GetTask(ctx context.Context, task *ssov1.GetTaskRequest) (*ssov1.GetTaskResponse, error)
	UpdateTask(ctx context.Context, task *ssov1.UpdateTaskRequest) (*ssov1.UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, task *ssov1.DeleteTaskRequest) (*ssov1.DeleteTaskResponse, error)
}
