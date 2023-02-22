package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var Redis *redis.Client

func getConf() string {
	dsn := os.Getenv("REDIS_DSN")
	if dsn != "" {
		return dsn
	}
	return "redis://localhost:6379" // test env
}

func Connect() {
	opt, err := redis.ParseURL(getConf())
	if err != nil {
		log.Fatal("Redis connection:", err)
	}
	Redis = redis.NewClient(opt)
	ping := Redis.Ping(context.Background())
	if ping.Val() != "PONG" {
		log.Fatal("Failed to connect redis")
	}
}
