package user

import (
	"my_project/internal/models"
	respository "my_project/internal/respository"
	"time"
)

type TaskService struct {
	userRepo *respository.UserRepository
	taskRepo *respository.TaskRepository
}

func NewTaskService() *TaskService {
	return &TaskService{
		userRepo: respository.NewUserRepository(),
		taskRepo: respository.NewTaskRepository(),
	}
}

func (t *TaskService) CreateATaskByUser(title, description, state string, complete bool, startDate, dueDate string, userID int) (*models.Task, error) {
	const layout = "2006-01-02" // or your desired date format
	parsedStartDate, errStartDate := time.Parse(layout, startDate)
	if errStartDate != nil {
		return nil, errStartDate
	}

	parsedDueDate, errDueDate := time.Parse(layout, dueDate)
	if errDueDate != nil {
		return nil, errDueDate
	}
	task, err := t.taskRepo.InsertATask(title, description, state, complete, parsedStartDate, parsedDueDate, userID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskService) UpdateATaskByUser(taskID int, title, description, state string, complete bool, startDate, dueDate string, userID int) (*models.Task, error) {
	const layout = "2006-01-02"

	parsedStartDate, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}

	parsedDueDate, err := time.Parse(layout, dueDate)
	if err != nil {
		return nil, err
	}

	task, err := t.taskRepo.UpdateTask(taskID, title, description, state, complete, parsedStartDate, parsedDueDate, userID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskService) DeleteATaskByUser(taskID int, userID int) error {
	// Repo sẽ đảm bảo chỉ xóa task thuộc user này
	err := t.taskRepo.DeleteATask(taskID, userID)
	if err != nil {
		return err
	}
	return nil
}
