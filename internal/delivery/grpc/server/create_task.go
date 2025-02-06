package grpc_server

import (
	"context"

	"github.com/google/uuid"
	ssov1 "github.com/nottee-project/protos/fleap_protos"
	models "github.com/nottee-project/task_service/internal/models/task"
	"github.com/nottee-project/task_service/internal/service/task"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewTaskService(repo task.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, req *ssov1.CreateTaskRequest) (*ssov1.CreateTaskResponse, error) {
	// Преобразуем gRPC-запрос в параметры, ожидаемые TaskRepository
	taskParams := models.CreateTaskParams{
		Title: req.Title,
		Body:  req.Body,
	}

	// Вызываем бизнес-логику
	createdTask, err := s.repo.CreateTask(ctx, taskParams)
	if err != nil {
		return nil, err
	}
	// Преобразуем результат в gRPC-ответ
	return &ssov1.CreateTaskResponse{
		Task: &ssov1.Task{
			Id:        createdTask.Id,
			UserId:    uuid.UUID.String(createdTask.UserId),
			Title:     createdTask.Title,
			Body:      createdTask.Body,
			CreatedAt: timestamppb.New(createdTask.CreatedAt),
		},
	}, nil
}
