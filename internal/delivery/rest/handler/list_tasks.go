package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/task_service/internal/models/task"
)

func (t *TaskHandler) ListTasks(c echo.Context) error {
	userId, ok := c.Get("user_id").(uuid.UUID)
	if !ok || userId == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	params := models.ListTasksParams{
		UserId: userId,
	}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	tasks, err := t.TaskSrv.ListTasks(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve tasks",
		})
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return c.Render(http.StatusOK, "tasks.html", tasks)
	}

	return c.JSON(http.StatusOK, tasks)
}

// func (t *TaskHandler) Index(c echo.Context) error {
// 	params := models.ListTasksParams{
// 		UserId: c.Get("user_id").(string),
// 	}

// 	tasks, err := t.TaskSrv.ListTasks(context.Background(), params)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to retrieve tasks",
// 		})
// 	}

// 	return c.Render(http.StatusOK, "index.html", tasks)
// }
