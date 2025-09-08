package models

import (
	"errors"
	"my_project/internal/helpers"

	"gorm.io/gorm"
)

type Account struct {
	ID       int    `gorm:"primaryKey"`
	UserID   int    `gorm:"uniqueIndex;not null"` // mỗi user có 1 account
	Username string `gorm:"uniqueIndex;size:50;not null"`
	Password string `gorm:"size:255;not null"`       // đã hash (không lưu plain text!)
	Provider string `gorm:"size:20;default:'local'"` // local, google, facebook...
	Timestamp

	// Quan hệ 1-1: 1 account ↔ 1 user
	User User `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID;references:ID"`
}

// BeforeCreate: hash password bằng helper
func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate password
	if !helpers.ValidatePassword(a.Password) {
		return errors.New("password must be at least 8 characters")
	}
	hashed, err := helpers.HashPassword(a.Password)
	if err != nil {
		return err
	}
	a.Password = hashed
	return
}

// BeforeUpdate: hash lại nếu password thay đổi
func (a *Account) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") {
		// Validate password
		if !helpers.ValidatePassword(a.Password) {
			return errors.New("password must be at least 8 characters")
		}

		hashed, err := helpers.HashPassword(a.Password)
		if err != nil {
			return err
		}
		a.Password = hashed
	}
	return
}
