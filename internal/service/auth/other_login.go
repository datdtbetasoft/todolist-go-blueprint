package auth

import (
	"context"
	"fmt"
	"my_project/internal/constants"
	redis_server "my_project/internal/database"

	"github.com/go-redis/redis/v8"
)

var keyPrefix = "Login_user_"

// SetLoginStatus Save login state in redis
func SetLoginStatus(userId string, sessionId string) error {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := keyPrefix + userId

	err := client.Set(context.Background(), key, sessionId, ttl).Err()

	return err
}

// CheckLoginSession Check login session available
func CheckLoginSession(userId, requestToken string) (bool, error) {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := keyPrefix + userId

	value, err := client.Get(context.Background(), key).Result()
	// There was no value, so no login state
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if value != requestToken {
		return false, err
	}
	return true, nil
}

// ExistOtherLogin login exists in another session
func ExistOtherLogin(userId string) (bool, string, error) {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := keyPrefix + userId

	// check who is currently logged in
	// key: userId
	// value: sessionID
	loginedSessionID, err := client.Get(context.Background(), key).Result()

	// There was no value, so no login state
	if err == redis.Nil {
		return false, "", nil
	}
	if err != nil {
		return false, "", err
	}

	return true, loginedSessionID, nil

}

// LogoutForOther Delete the first logged-in session and update with the last win
func LogoutForOther(userId string, sessionId string, deleteTargetSessionId string) error {
	// TODO: transaction
	err := SetLoginStatus(userId, sessionId)
	if err != nil {
		fmt.Println("error: SetLoginStatus", err)
		return err
	}

	// remove the session id that already exists
	// session ID management
	errDelete := Delete(deleteTargetSessionId)
	if errDelete != nil {
		fmt.Println("error: delete sessionID", errDelete)
		return errDelete
	}

	return nil
}

// DeleteLoginStatus delete record from login status management table
func DeleteLoginStatus(userId string) error {
	client := redis_server.GetRedisServer(constants.Authentication)
	key := keyPrefix + userId

	err := client.Del(context.Background(), key).Err()
	if err != nil {
		fmt.Println("redis:1 delete", err)
		return err
	}
	return nil
}
