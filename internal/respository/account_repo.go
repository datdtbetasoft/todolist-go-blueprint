package repository

import (
	"my_project/internal/models"

	"gorm.io/gorm"
)

type AccRepository struct {
	*BaseRepository
}

func NewAccRepository() *AccRepository {
	return &AccRepository{BaseRepository: baseRepo}
}

func (r *AccRepository) InsertAAcc(user models.User, password string, provider string) (*models.Account, error) {
	account := &models.Account{
		Username: *user.Email,
		Password: password,
		Provider: provider,
		User:     user,
	}

	if err := r.db.Create(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccRepository) FindByUsername(username string) (*models.Account, error) {
	var account models.Account

	// find account by username
	if err := r.db.Preload("User").Where("username = ?", username).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccRepository) WithTx(tx *gorm.DB) *AccRepository {
	return &AccRepository{
		BaseRepository: &BaseRepository{db: tx},
	}
}
