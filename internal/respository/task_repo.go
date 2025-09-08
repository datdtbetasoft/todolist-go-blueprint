package repository

import (
	"errors"
	"fmt"
	"my_project/internal/constants"
	"my_project/internal/models"
	"time"
)

type TaskRepository struct {
	*BaseRepository
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{BaseRepository: baseRepo}
}

func (r *TaskRepository) InsertATask(title, description, state string, complete bool, startDate, dueDate time.Time, userID int) (*models.Task, error) {
	task := &models.Task{
		Title:       title,
		Description: description,
		DueDate:     &dueDate,
		StartDate:   &startDate,
		State:       constants.TaskState(state),
		Completed:   complete,
		UserID:      userID,
	}

	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTask(taskID int, title, description, state string, complete bool, startDate, dueDate time.Time, userID int) (*models.Task, error) {
	var task models.Task
	// Tìm task theo ID và UserID (để chắc chắn user này sở hữu task)
	if err := r.db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		return nil, errors.New("task not found or you do not have permission to update it")
	}

	// Cập nhật dữ liệu
	task.Title = title
	task.Description = description
	task.State = constants.TaskState(state)
	task.Completed = complete
	task.StartDate = &startDate
	task.DueDate = &dueDate

	if err := r.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) DeleteATask(taskID int, userID int) error {
	// Xóa task dựa theo ID và userID để tránh xóa nhầm task của user khác
	result := r.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("task not found or not belong to user")
	}
	return nil
}
