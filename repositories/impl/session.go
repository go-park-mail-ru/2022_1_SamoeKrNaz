package impl

import (
	"PLANEXA_backend/auth_microservice/server/handler"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"context"
)

const (
	CookieTime = 259200 // 3 суток
)

type SessionRepositoryImpl struct {
	client handler.AuthCheckerClient
	ctx    context.Context
}

func CreateRepo(cl handler.AuthCheckerClient) repositories.SessionRepository {
	return &SessionRepositoryImpl{client: cl, ctx: context.Background()}
}

func (redisConnect SessionRepositoryImpl) SetSession(session models.Session) error {
	_, err := redisConnect.client.Create(redisConnect.ctx, &handler.SessionModel{SESSIONVALUE: session.CookieValue, USERID: uint64(session.UserId)})
	return err
}

func (redisConnect SessionRepositoryImpl) GetSession(cookieValue string) (uint64, error) {
	userId, err := redisConnect.client.Get(redisConnect.ctx, &handler.SessionValue{Value: cookieValue})
	return userId.ID, err
}

func (redisConnect SessionRepositoryImpl) DeleteSession(cookieValue string) error {
	_, err := redisConnect.client.Delete(redisConnect.ctx, &handler.SessionValue{Value: cookieValue})
	return err
}
