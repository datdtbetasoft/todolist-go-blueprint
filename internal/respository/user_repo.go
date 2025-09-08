package repository

import (
	"my_project/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{BaseRepository: baseRepo}
}

func (r *UserRepository) InsertAUser(name, email string, birthday time.Time) (*models.User, error) {
	user := &models.User{
		Name:     name,
		Email:    &email,
		Birthday: &birthday,
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: &BaseRepository{db: tx},
	}
}
