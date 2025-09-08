package auth

import (
	"context"
	"errors"
	"fmt"
	"my_project/internal/constants"
	"time"

	redis_server "my_project/internal/database"

	"github.com/go-redis/redis/v8"
)

type SessionRedis struct{}

var (
	ttl      = 24 * time.Hour // 24h
	ttl_code = 15 * time.Minute
)

const verifyCodePrefix = "Verify_Code_"
const registerCodePrefix = "Register_Code_"

// SetVerifyCodeStatus Save verify code state in redis
func SetVerifyCodeStatus(userId string, code string) error {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := verifyCodePrefix + userId

	err := client.Set(context.Background(), key, code, ttl_code).Err()

	return err
}

// RemoveVerifyCodeStatus Remove verify code state in redis
func RemoveVerifyCodeStatus(userId string) error {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := verifyCodePrefix + userId

	err := client.Del(context.Background(), key).Err()
	if err != nil {
		fmt.Println("redis:1 delete", err)
		return err
	}

	return nil
}

// CheckVerifyCodeStatus Check verify code state in redis
func CheckVerifyCodeStatus(userId, verifyCode string) (bool, error) {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := verifyCodePrefix + userId

	result, err := client.Get(context.Background(), key).Result()
	// There was no value, so no login state
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if result != verifyCode {
		return false, nil
	}

	return true, nil
}

// SetRegisterCodeStatus Save registration verification code in redis using email as key
func SetRegisterCodeStatus(email string, code string) error {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := registerCodePrefix + email

	err := client.Set(context.Background(), key, code, ttl_code).Err()

	return err
}

// RemoveRegisterCodeStatus Remove registration verification code from redis
func RemoveRegisterCodeStatus(email string) error {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := registerCodePrefix + email

	err := client.Del(context.Background(), key).Err()
	if err != nil {
		fmt.Println("redis:1 delete register code", err)
		return err
	}

	return nil
}

// CheckRegisterCodeStatus Check if registration verification code is valid
func CheckRegisterCodeStatus(email, verifyCode string) (bool, error) {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := registerCodePrefix + email

	result, err := client.Get(context.Background(), key).Result()
	// There was no value, so no verification code
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if result != verifyCode {
		return false, nil
	}

	return true, nil
}

func Update(sessionID string) error {
	client := redis_server.GetRedisServer(constants.Authentication)

	err := client.Expire(context.Background(), sessionID, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func Create(sessionID string, val string) error {
	client := redis_server.GetRedisServer(constants.Authentication)
	err := client.SetNX(context.Background(), sessionID, val, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(sessionID string) (string, error) {
	client := redis_server.GetRedisServer(constants.Authentication)

	val, err := client.Get(context.Background(), sessionID).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Delete(sessionID string) error {
	client := redis_server.GetRedisServer(constants.Authentication)

	err := client.Del(context.Background(), sessionID).Err()
	if err != nil {
		return err
	}
	return nil
}
