package grpc_server

import (
	"context"

	"github.com/google/uuid"
	ssov1 "github.com/nottee-project/protos/fleap_protos"
	models "github.com/nottee-project/task_service/internal/models/task"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *TaskService) ListTasks(ctx context.Context, req *ssov1.ListTasksRequest) (*ssov1.ListTasksResponse, error) {
	// Получаем список задач из репозитория
	tasks, err := s.repo.ListTasks(ctx, models.ListTasksParams{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	// Конвертируем список задач в gRPC-формат
	var grpcTasks []*ssov1.Task
	for _, t := range tasks {
		grpcTasks = append(grpcTasks, &ssov1.Task{
			Id:        t.ID,
			UserId:    uuid.UUID.String(t.UserID),
			Title:     t.Title,
			Body:      t.Body,
			CreatedAt: timestamppb.New(t.CreatedAt),
		})
	}

	return &ssov1.ListTasksResponse{Tasks: grpcTasks}, nil
}
