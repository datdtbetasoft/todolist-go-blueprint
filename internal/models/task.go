package models

import (
	"my_project/internal/constants"
	"time"
)

type TodoState string

type Task struct {
	ID          int                 `gorm:"primaryKey"`
	UserID      int                 `gorm:"index;not null"`    // foreign key -> users.id
	Title       string              `gorm:"size:255;not null"` // tiêu đề task
	Description string              `gorm:"type:text"`         // mô tả chi tiết
	Completed   bool                `gorm:"default:false"`     // trạng thái hoàn thành
	State       constants.TaskState `gorm:"type:varchar(20);default:'pending'"`

	StartDate *time.Time // ngày bắt đầu (có thể null)
	DueDate   *time.Time // hạn chót (có thể null)

	User User `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID;references:ID"`

	Timestamp
}
