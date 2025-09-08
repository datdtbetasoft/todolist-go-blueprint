package repository

import (
	postgres "my_project/internal/database"

	"gorm.io/gorm"
)

var baseRepo = &BaseRepository{
	db: postgres.GetDB(),
}

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository() *BaseRepository {
	db := postgres.GetDB()
	return &BaseRepository{db: db}
}
