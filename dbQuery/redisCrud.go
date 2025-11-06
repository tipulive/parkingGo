package dbQuery

import (
	"log"
	"parking/db"
)

func CreateRedis() {

	// Set value
	err := db.RDS.Set(db.CTXRedis, "key1", "Hello from main", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

}
func ReadRedis() {
	// Get value
	val, err := db.RDS.Get(db.CTXRedis, "key1").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("key1", val)
}
func DeleteRedis() {
	// Delete value
	err := db.RDS.Del(db.CTXRedis, "key1").Err()
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateRedis() {
	// Update value
	err := db.RDS.Set(db.CTXRedis, "key1", "Updated Value", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}
func ExistsRedis() {
	// Check if key exists
	val, err := db.RDS.Exists(db.CTXRedis, "key1").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("key1 exists:", val > 0)
}
func ExpireRedis() {
	// Set expiration
	err := db.RDS.Expire(db.CTXRedis, "key1", 60*5).Err() // 5 minutes
	if err != nil {
		log.Fatal(err)
	}
}
func TTLRedis() {
	// Get TTL
	val, err := db.RDS.TTL(db.CTXRedis, "key1").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("key1 TTL:", val)
}
func FlushRedis() {
	// Flush all keys
	err := db.RDS.FlushAll(db.CTXRedis).Err()
	if err != nil {
		log.Fatal(err)
	}
}
func PingRedis() {
	// Ping Redis
	pong, err := db.RDS.Ping(db.CTXRedis).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Redis Ping:", pong)
}
func TransactionRedis() {
	// Start transaction
	pipe := db.RDS.TxPipeline()

	// Queue commands
	pipe.Set(db.CTXRedis, "key1", "value1", 0)
	pipe.Set(db.CTXRedis, "key2", "value2", 0)

	// Execute transaction
	_, err := pipe.Exec(db.CTXRedis)
	if err != nil {
		log.Fatal(err)
	}
}
func LuaScriptRedis() {
	// Lua script
	script := `
		return redis.call("SET", KEYS[1], ARGV[1])
	`
	// Execute script
	result, err := db.RDS.Eval(db.CTXRedis, script, []string{"key1"}, "value1").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Lua Script Result:", result)
}
func PipelineRedis() {
	// Create pipeline
	pipe := db.RDS.Pipeline()

	// Queue commands
	pipe.Set(db.CTXRedis, "key1", "value1", 0)
	pipe.Set(db.CTXRedis, "key2", "value2", 0)

	// Execute pipeline
	_, err := pipe.Exec(db.CTXRedis)
	if err != nil {
		log.Fatal(err)
	}
}
