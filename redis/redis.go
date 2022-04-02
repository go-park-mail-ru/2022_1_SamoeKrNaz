package main

import (
	"PLANEXA_backend/handlers"
	"PLANEXA_backend/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type RedisConnect struct {
	client *redis.Client
}

func (redisConnect RedisConnect) ConnectToRedis() *RedisConnect {
	return &RedisConnect{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
}

func (redisConnect RedisConnect) SetCookie(session models.Session) error {
	return redisConnect.client.Set(session.CookieValue, strconv.Itoa(int(session.UserId)), time.Duration(handlers.CookieTime)).Err()
}

func (redisConnect RedisConnect) GetCookie(session models.Session) (uint64, error) {
	value, err := redisConnect.client.Get(session.CookieValue).Uint64()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (redisConnect RedisConnect) DeleteCookie(session models.Session) error {
	err := redisConnect.client.Del(session.CookieValue).Err()
	if err != nil {
		return err
	}
	return nil
}
