package db

import (
	//"database/sql"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"parking/config" // replace with your actual module path

	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

//var DB *sql.DB

var (
	DB       *sqlx.DB
	RDS      *redis.Client
	CTXRedis = context.Background()
)

func InitDB(cfg config.MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	test := os.Getenv("DB_USER")
	var err error
	//DB, err = sql.Open("mysql", dsn)
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	fmt.Println("Database connected!%v", test)
}
func InitRedis() {
	RDS = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379", // Redis address
		Password:     "",               // No password set
		DB:           0,                // Default DB
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})

	// Test connection
	pong, err := RDS.Ping(CTXRedis).Result()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
		return
	}

	fmt.Println("✅ Redis connected:", pong)
}
