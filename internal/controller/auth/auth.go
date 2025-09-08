package auth

import (
	"errors"
	"my_project/internal/helpers"
	respository "my_project/internal/respository"
	"strconv"
)

type AuthController struct {
	userRepo    *respository.UserRepository
	accountRepo *respository.AccRepository
}

func NewAuthController() *AuthController {
	return &AuthController{
		userRepo:    respository.NewUserRepository(),
		accountRepo: respository.NewAccRepository(),
	}
}

func (s *AuthController) Login(username string, password string) (string, error) {
	acc, err := s.accountRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if acc == nil {
		// user not found
		return "", errors.New("user not found")
	}

	if !helpers.CheckPasswordHash(password, acc.Password) {
		// password not match
		return "", errors.New("password wrong")
	}

	token, errGenToken := GenerateToken(strconv.Itoa(int(acc.UserID)))
	if errGenToken != nil {
		return "", errGenToken
	}
	// userJSON, errJson := json.Marshal(acc.User)
	// if errJson != nil {
	// 	return "", errJson
	// }
	sessionId := "Login_user_" + strconv.Itoa(int(acc.ID))
	tokenOld, errGetOld := Get(sessionId)
	if errGetOld != nil {
		return "", errGenToken
	}
	if tokenOld == "" {
		Create(sessionId, token)
		return token, nil
	}
	return tokenOld, nil
}
