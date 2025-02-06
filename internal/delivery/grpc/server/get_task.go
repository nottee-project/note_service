package grpc_server

import (
	"context"

	"github.com/google/uuid"
	ssov1 "github.com/nottee-project/protos/fleap_protos"
)

func (s *serverAPI) GetTask(ctx context.Context, req *ssov1.GetTaskRequest) (*ssov1.GetTaskResponse, error) {
	// Получаем задачу из репозитория
	task, err := s.repo.GetTask(ctx, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}

	// Возвращаем gRPC-ответ
	return &ssov1.GetTaskResponse{
			Id:     task.ID,
			UserId: uuid.UUID.String(task.UserID),
			Title:  task.Title,
			Body:   task.Body,
		},
		nil
}
