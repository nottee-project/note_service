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
}

func Register(gRPCServer *grpc.Server, taskService TaskService) {
	ssov1.RegisterTaskServiceServer(gRPCServer, &serverAPI{taskService: taskService})
	log.Println("TaskService gRPC server registered")
}



func (s *serverAPI) CreateTask(ctx context.Context, req *ssov1.CreateTaskRequest) (*ssov1.CreateTaskResponse, error) {
	log.Println("CreateTask called with:", req)
	return s.taskService.CreateTask(ctx, req)
}

func (s *serverAPI) ListTasks(ctx context.Context, req *ssov1.ListTasksRequest) (*ssov1.ListTasksResponse, error) {
	log.Println("ListTasks called")
	return s.taskService.ListTasks(ctx, req)
}

func (s *serverAPI) GetTask(ctx context.Context, req *ssov1.GetTaskRequest) (*ssov1.GetTaskResponse, error) {
	log.Println("GetTask called with:", req)
	return s.taskService.GetTask(ctx, req)
}

func (s *serverAPI) UpdateTask(ctx context.Context, req *ssov1.UpdateTaskRequest) (*ssov1.UpdateTaskResponse, error) {
	log.Println("UpdateTask called with:", req)
	return s.taskService.UpdateTask(ctx, req)
}

func (s *serverAPI) DeleteTask(ctx context.Context, req *ssov1.DeleteTaskRequest) (*ssov1.DeleteTaskResponse, error) {
	log.Println("DeleteTask called with:", req)
	return s.taskService.DeleteTask(ctx, req)
}


