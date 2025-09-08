package user

import (
	"my_project/internal/models"
	respository "my_project/internal/respository"
	"time"

	postgres "my_project/internal/database"

	"gorm.io/gorm"
)

type UserController struct {
	db          *gorm.DB
	userRepo    *respository.UserRepository
	accountRepo *respository.AccRepository
}

func NewUserController() *UserController {
	return &UserController{
		db:          postgres.GetDB(),
		userRepo:    respository.NewUserRepository(),
		accountRepo: respository.NewAccRepository(),
	}
}

func (s *UserController) Register(name, email string, password string, birthday time.Time, provider string) (*models.User, *models.Account, error) {
	var user *models.User
	var account *models.Account

	err := s.db.Transaction(func(tx *gorm.DB) error {
		uRepo := s.userRepo.WithTx(tx)
		aRepo := s.accountRepo.WithTx(tx)

		var err error
		user, err = uRepo.InsertAUser(name, email, birthday)
		if err != nil {
			return err
		}

		account, err = aRepo.InsertAAcc(*user, password, provider)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}
	return user, account, nil
}
