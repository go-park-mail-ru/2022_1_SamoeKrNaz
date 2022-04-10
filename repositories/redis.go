package repositories

import (
	"PLANEXA_backend/models"
	"github.com/go-redis/redis"
	"time"
)

const (
	CookieTime = 259200 // 3 суток
)

type RedisRepository struct {
	client *redis.Client
}

func ConnectToRedis() *RedisRepository {
	return &RedisRepository{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
}

func (redisConnect RedisRepository) SetSession(session models.Session) error {
	return redisConnect.client.Set(session.CookieValue, session.UserId, time.Duration(CookieTime)).Err()
}

func (redisConnect RedisRepository) GetSession(cookieValue string) (uint64, error) {
	value, err := redisConnect.client.Get(cookieValue).Uint64()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (redisConnect RedisRepository) DeleteSession(cookieValue string) error {
	err := redisConnect.client.Del(cookieValue).Err()
	if err != nil {
		return err
	}
	return nil
}
