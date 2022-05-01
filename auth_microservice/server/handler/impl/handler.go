package impl

import (
	"PLANEXA_backend/auth_microservice/server/handler"
	"PLANEXA_backend/auth_microservice/server/usecase"
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"context"
)

type SessionServerImpl struct {
	sessUseCase usecase.SessionUseCase
	handler.UnimplementedAuthCheckerServer
}

func CreateSessionServer(sessUseCase usecase.SessionUseCase) handler.AuthCheckerServer {
	return &SessionServerImpl{sessUseCase: sessUseCase}
}

func (sessServer *SessionServerImpl) Create(ctx context.Context, in *handler.SessionModel) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	sess := models.Session{UserId: uint(in.USERID), CookieValue: in.SESSIONVALUE}

	err := sessServer.sessUseCase.SetSession(sess)
	return &handler.Nothing{}, err
}

func (sessServer *SessionServerImpl) Get(ctx context.Context, in *handler.SessionValue) (*handler.SessionID, error) {
	if in == nil {
		return &handler.SessionID{ID: 0}, customErrors.ErrBadInputData
	}

	userId, err := sessServer.sessUseCase.GetSession(in.Value)
	return &handler.SessionID{ID: userId}, err
}

func (sessServer *SessionServerImpl) Delete(ctx context.Context, in *handler.SessionValue) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}

	err := sessServer.sessUseCase.DeleteSession(in.Value)
	return &handler.Nothing{}, err
}
