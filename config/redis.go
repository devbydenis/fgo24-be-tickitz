package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)
func RedisConnect() *redis.Client {
    godotenv.Load()
    
    db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
    if err != nil {
        fmt.Println("Failed to convert DB to int:", err)
    }

    var RedisClient = redis.NewClient(&redis.Options{
        Addr:	  os.Getenv("REDIS_ADDRESS"),
        Password: os.Getenv("REDIS_PASSWORD"), // No password set
        DB:		  db,  // Use default DB
    })

    return RedisClient
}



