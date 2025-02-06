package grpc_server

import (
	"context"

	ssov1 "github.com/nottee-project/protos/fleap_protos"
)

func (s *TaskService) DeleteTask(ctx context.Context, req *ssov1.DeleteTaskRequest) (*ssov1.DeleteTaskResponse, error) {
	// Удаляем задачу из репозитория
	err := s.repo.DeleteTask(ctx, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}

	return &ssov1.DeleteTaskResponse{Message: "Task deleted"}, nil
}
