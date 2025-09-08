package controller

import (
	"my_project/internal/https/request"
	"my_project/internal/https/response"
	"net/http"
	"strconv"

	taskServ "my_project/internal/service/task"

	"github.com/labstack/echo/v4"
)

type TaskController struct{}

func NewTaskController() *TaskController {
	return new(TaskController)
}

// Add user-related handler methods here
func (ctl *TaskController) CreateATask(c echo.Context) error {
	userId := c.Get("userId").(int)
	params, err := BindAndValidate[request.CreateTaskRequest](c)
	if err != nil {
		return err
	}

	task, errServ := taskServ.NewTaskService().CreateATaskByUser(params.Title, params.Description, params.State, params.Completed, params.StartDate, params.DueDate, userId)

	if errServ != nil {
		return c.JSON(http.StatusInternalServerError, response.NewResponse(
			"500",
			"Failed to create task",
			errServ.Error(),
		))
	}

	return c.JSON(http.StatusOK, response.NewResponse(
		"200",
		"Login successful",
		task,
	))
}

func (ctl *TaskController) UpdateATask(c echo.Context) error {
	// Lấy taskID từ URL param
	taskIDParam := c.Param("id")
	taskID, err := strconv.ParseInt(taskIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(
			"400",
			"Invalid task ID",
			err.Error(),
		))
	}

	// Lấy userID từ middleware (JWT gán vào context)
	userID := c.Get("userId").(int)

	// Bind và validate request body
	params, err := BindAndValidate[request.UpdateTaskRequest](c)
	if err != nil {
		return err
	}

	// Gọi service để update
	task, errServ := taskServ.NewTaskService().UpdateATaskByUser(
		int(taskID),
		params.Title,
		params.Description,
		params.State,
		params.Completed,
		params.StartDate,
		params.DueDate,
		userID,
	)

	if errServ != nil {
		return c.JSON(http.StatusInternalServerError, response.NewResponse(
			"500",
			"Failed to update task",
			errServ.Error(),
		))
	}

	return c.JSON(http.StatusOK, response.NewResponse(
		"200",
		"Task updated successfully",
		task,
	))
}

func (ctl *TaskController) DeleteATask(c echo.Context) error {
	// Lấy task ID từ param
	taskIDParam := c.Param("id")
	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(
			"400",
			"Invalid task ID",
			nil,
		))
	}

	// Lấy userId từ context (được middleware gắn vào)
	userId := c.Get("userId").(int)

	// Gọi service
	err = taskServ.NewTaskService().DeleteATaskByUser(taskID, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewResponse(
			"500",
			"Failed to delete task",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, response.NewResponse(
		"200",
		"Task deleted successfully",
		nil,
	))
}
