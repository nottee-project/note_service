package handler

import (
	"github.com/nottee-project/task_service/internal/service/task"
)

type TaskHandler struct {
	TaskSrv *task.TaskService
}

func NewTaskHandler(TaskSrv *task.TaskService) TaskHandler {
	return TaskHandler{
		TaskSrv: TaskSrv,
	}
}
