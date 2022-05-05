package impl

import (
	"PLANEXA_backend/auth_microservice/server_session_ms/handler"
	"PLANEXA_backend/auth_microservice/server_session_ms/usecase"
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/metrics"
	"PLANEXA_backend/models"
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

type SessionServerImpl struct {
	sessUseCase usecase.SessionUseCase
	handler.UnimplementedAuthCheckerServer
}

func CreateSessionServer(sessUseCase usecase.SessionUseCase) handler.AuthCheckerServer {
	return &SessionServerImpl{sessUseCase: sessUseCase}
}

func (sessServer *SessionServerImpl) Create(ctx context.Context, in *handler.SessionModel) (*handler.Nothing, error) {
	timer := prometheus.NewTimer(metrics.DurationSession.WithLabelValues("create"))
	if in == nil {
		metrics.Session.WithLabelValues("500", "nil \"in\" create session").Inc()
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	sess := models.Session{UserId: uint(in.USERID), CookieValue: in.SESSIONVALUE}

	err := sessServer.sessUseCase.SetSession(sess)
	if err != nil {
		metrics.Session.WithLabelValues("500", "error in create session").Inc()
		return &handler.Nothing{}, err
	}
	metrics.Session.WithLabelValues("200", "success in create session").Inc()
	timer.ObserveDuration()
	return &handler.Nothing{}, nil
}

func (sessServer *SessionServerImpl) Get(ctx context.Context, in *handler.SessionValue) (*handler.SessionID, error) {
	timer := prometheus.NewTimer(metrics.DurationSession.WithLabelValues("get"))
	if in == nil {
		metrics.Session.WithLabelValues("500", "nil \"in\", get session").Inc()
		return &handler.SessionID{ID: 0}, customErrors.ErrBadInputData
	}

	userId, err := sessServer.sessUseCase.GetSession(in.Value)
	if err != nil {
		metrics.Session.WithLabelValues("500", "error in get session").Inc()
		return &handler.SessionID{ID: 0}, err
	}
	metrics.Session.WithLabelValues("200", "success in get session").Inc()
	timer.ObserveDuration()
	return &handler.SessionID{ID: userId}, nil
}

func (sessServer *SessionServerImpl) Delete(ctx context.Context, in *handler.SessionValue) (*handler.Nothing, error) {
	timer := prometheus.NewTimer(metrics.DurationSession.WithLabelValues("delete"))
	if in == nil {
		metrics.Session.WithLabelValues("500", "nil \"in\" delete session").Inc()
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	err := sessServer.sessUseCase.DeleteSession(in.Value)
	if err != nil {
		metrics.Session.WithLabelValues("500", "error in delete session").Inc()
		return &handler.Nothing{}, err
	}
	metrics.Session.WithLabelValues("200", "success in delete session").Inc()
	timer.ObserveDuration()
	return &handler.Nothing{}, nil
}
