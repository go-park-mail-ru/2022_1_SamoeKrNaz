package planexa_redis

import (
	"PLANEXA_backend/handlers"
	"PLANEXA_backend/models"
	"github.com/go-redis/redis"
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

func (redisConnect RedisConnect) SetSession(session models.Session) error {
	return redisConnect.client.Set(session.CookieValue, session.UserId, time.Duration(handlers.CookieTime)).Err()
}

func (redisConnect RedisConnect) GetSession(cookieValue string) (uint64, error) {
	value, err := redisConnect.client.Get(cookieValue).Uint64()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (redisConnect RedisConnect) DeleteSession(cookieValue string) error {
	err := redisConnect.client.Del(cookieValue).Err()
	if err != nil {
		return err
	}
	return nil
}
