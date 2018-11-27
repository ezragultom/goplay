package conn

import (
	"time"

	"gopkg.in/redis.v5"
)

var redisClient map[string]*redis.Client = map[string]*redis.Client{}

func RedisClient(node string) *redis.Client {

	if val, ok := redisClient[node]; ok {
		return val
	}

	host := "redis:6379"
	password := ""
	maxRetries := 0
	db := 0
	idleTimeout := time.Duration(0)

	client := redis.NewClient(&redis.Options{
		Addr:        host,
		Password:    password,
		DB:          db,
		MaxRetries:  maxRetries,
		IdleTimeout: time.Second * idleTimeout,
	})

	redisClient[node] = client

	return client
}
