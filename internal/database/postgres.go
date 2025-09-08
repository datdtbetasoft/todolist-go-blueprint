package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"my_project/internal/config"
	"my_project/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB
var once sync.Once

// migrate tự động tạo bảng từ models
func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Task{},
	)
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)
	db, err := postgresDB.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		return stats
	}

	// Ping database
	if err := db.PingContext(ctx); err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		return stats
	}

	// Database is up
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Connection pool statistics
	dbStats := db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate load
	if dbStats.OpenConnections > 40 {
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}
	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}
	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

func ClosePostgres() error {
	log.Println("Disconnected from database")
	db, err := postgresDB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func InitPostgres(ctx context.Context) {
	c := config.GetConfig()
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
			c.GetString("db.postgresql.host"),
			c.GetString("db.postgresql.username"),
			c.GetString("db.postgresql.password"),
			c.GetString("db.postgresql.database"),
			c.GetString("db.postgresql.port"),
			c.GetString("db.postgresql.schema"),
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Failed to connect database: %v", err)
		}

		// Gán vào biến global
		postgresDB = conn

		// Lấy sql.DB để config pool
		sqlDB, err := conn.DB()
		if err != nil {
			log.Fatalf("❌ Failed to get sql.DB: %v", err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		sqlDB.SetConnMaxIdleTime(10 * time.Minute)

		log.Println("✅ Connected to database")

		// Auto migrate
		if err := migrate(conn); err != nil {
			log.Fatalf("❌ Migration failed: %v", err)
		}
		log.Println("✅ Auto migration done")
	})
}

func GetDB() *gorm.DB {
	return postgresDB
}
