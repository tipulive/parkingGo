package db

import (
	//"database/sql"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"parking/config" // replace with your actual module path

	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

// var DB *sql.DB
const (
	RoleVersionKey    = "roles_version"
	RoleVersionPubSub = "roles_version_changed"
	CacheTTL          = 5 * time.Second
	MaxBackoffRetries = 5
	AccessTokenTTL    = 15 * time.Minute
	RefreshTokenTTL   = 7 * 24 * time.Hour // 7 days
)

var (
	DB           *sqlx.DB
	RDS          *redis.Client
	CTXRedis     = context.Background()
	RISSETOCACHE *ristretto.Cache // Local in-memory cache

)

func InitDB(cfg config.MySQLConfig) { //
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
func InitRedis() { //Redis
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

func init() { //Risseto
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e4,
		MaxCost:     1 << 20,
		BufferItems: 64,
	})
	if err != nil {
		log.Fatalf("failed to init cache: %v", err)
	}
	RISSETOCACHE = cache
}
func StartPubSubListener(ctx context.Context) {
	pubsub := RDS.Subscribe(ctx, RoleVersionPubSub)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Printf("pubsub subscribe error: %v", err)
		return
	}
	ch := pubsub.Channel()
	go func() {
		for msg := range ch {
			v, err := strconv.ParseUint(msg.Payload, 10, 64)
			if err != nil {
				continue
			}
			RISSETOCACHE.SetWithTTL(RoleVersionKey, v, 1, CacheTTL)
		}
	}()
}
