package handler

type TaskHandler struct {
	TaskSrv *task.TaskService
}

func NewTaskHandler(TaskSrv *task.TaskService) TaskHandler {
	return TaskHandler{
		TaskSrv: TaskSrv,
	}
}
