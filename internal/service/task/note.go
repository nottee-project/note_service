package task

type TaskService struct {
	TaskRepository
}

func New(repo TaskRepository) (*TaskService, error) {
	return &TaskService{TaskRepository: repo}, nil
}
