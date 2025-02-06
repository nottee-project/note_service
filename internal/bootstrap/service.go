package bootstrap

import (
	"github.com/nottee-project/task_service/internal/adapter/store"
	task_store "github.com/nottee-project/task_service/internal/adapter/store/task"
	"github.com/nottee-project/task_service/internal/config"
	"github.com/nottee-project/task_service/internal/service/task"
	"github.com/pkg/errors"
)

func CreateTaskService() (*task.TaskService, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, errors.Wrap(err, "CreateConfig")
	}

	dbStore, err := store.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	taskStore := &task_store.TaskStore{Store: dbStore}
	return task.New(taskStore)
}
