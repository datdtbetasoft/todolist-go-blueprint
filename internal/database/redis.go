package database

import (
	"fmt"
	"log"
	"my_project/internal/config"
	"my_project/internal/constants"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var redisServer map[int]*redis.Client

// Init initializes Redis connections for different DBs
func InitRedis() {
	redisServer = make(map[int]*redis.Client)
	cfg := config.GetConfig()

	host := cfg.GetString("redis.host")
	port := strconv.Itoa(cfg.GetInt("redis.port"))
	password := cfg.GetString("redis.password")

	// Danh sách DB cần khởi tạo
	dbList := map[int]int{
		constants.Authentication: constants.Authentication,
		// constants.Bubble:             constants.Bubble,
		// constants.Pin:                constants.Pin,
		constants.ShareLocationRedis: constants.Authentication,
		// constants.Chat:               constants.Authentication,
	}

	for key, dbIndex := range dbList {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       dbIndex,
		})

		// Ping để check connection
		if err := client.Ping(client.Context()).Err(); err != nil {
			log.Fatalf("❌ Failed to connect to Redis DB %d: %v", dbIndex, err)
		}

		log.Printf("✅ Connected to Redis DB %d at %s:%s", dbIndex, host, port)
		redisServer[key] = client
	}
}

// GetRedisServer returns Redis client by key
func GetRedisServer(key int) *redis.Client {
	return redisServer[key]
}

// Close closes all Redis connections
func CloseRedis() {
	for db, client := range redisServer {
		if client != nil {
			if err := client.Close(); err != nil {
				log.Printf("⚠️ Failed to close Redis DB %d: %v", db, err)
			} else {
				log.Printf("🔌 Closed Redis DB %d", db)
			}
		}
	}
}
