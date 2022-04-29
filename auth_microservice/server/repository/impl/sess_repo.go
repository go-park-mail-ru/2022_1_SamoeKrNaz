package repository_impl

import (
	"PLANEXA_backend/auth_microservice/server/repository"
	"PLANEXA_backend/models"
	"github.com/go-redis/redis"
)

const (
	CookieTime = 259200 // 3 суток
)

type SessionRepositoryImpl struct {
	client *redis.Client
}

func CreateSessRep() repository.SessionRedis {
	return &SessionRepositoryImpl{client: redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})}
}

func (redisConnect SessionRepositoryImpl) SetSession(session models.Session) error {
	return redisConnect.client.Do("SETEX", session.CookieValue, CookieTime, session.UserId).Err()
}

func (redisConnect SessionRepositoryImpl) GetSession(cookieValue string) (uint64, error) {
	value, err := redisConnect.client.Get(cookieValue).Uint64()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (redisConnect SessionRepositoryImpl) DeleteSession(cookieValue string) error {
	err := redisConnect.client.Del(cookieValue).Err()
	if err != nil {
		return err
	}
	return nil
}
