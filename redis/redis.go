package main

import (
	customErrors "PLANEXA_backend/errors"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type UserCookie struct {
	idU    uint
	cookie string
}

type RedisConnect struct {
	*redis.Client
}

func (redisConnect RedisConnect) ConnectToRedis() *RedisConnect {
	return &RedisConnect{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
}

func (redisConnect RedisConnect) SetCookie(user UserCookie) error {
	err := redisConnect.Set(strconv.Itoa(int(user.idU)), user.cookie, 259200*time.Second).Err()
	return err
}

func (redisConnect RedisConnect) GetCookie(user UserCookie) (string, error) {
	value, err := redisConnect.Get(strconv.Itoa(int(user.idU))).Result()
	if err != nil {
		return "", customErrors.ErrUnauthorized
	}
	return value, nil
}

func (redisConnect RedisConnect) DeleteCookie(user UserCookie) error {
	err := redisConnect.Del(strconv.Itoa(int(user.idU))).Err()
	if err != nil {
		return customErrors.ErrUserNotFound
	}
	return nil
}
