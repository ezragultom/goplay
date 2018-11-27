package conn

import (
	redis "gopkg.in/redis.v5"
)

type RedisHandler struct {
	client *redis.Client
}

func NewRedisHandler(c *redis.Client) *RedisHandler {
	return &RedisHandler{
		client: c,
	}
}

func (rh RedisHandler) Set(key, query string) (bool, error) {

	value, err := rh.client.HSet("ez_test_redis", key, query).Result()
	if err != nil {
		return value, err
	}
	return value, nil
}

func (rh RedisHandler) Get(key string) (string, error) {

	value, err := rh.client.HGet("ez_test_redis", key).Result()
	if err != nil {
		return value, err
	}
	return value, nil
}
