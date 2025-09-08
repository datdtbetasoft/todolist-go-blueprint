package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int        `gorm:"primaryKey"`
	Name     string     `gorm:"size:100;not null"`
	Email    *string    `gorm:"uniqueIndex;size:100"`
	Birthday *time.Time // A pointer to time.Time, can be null

	Tasks     []Task         `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Timestamp
}

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
