package usecase_impl

import (
	"PLANEXA_backend/auth_microservice/server/repository"
	"PLANEXA_backend/auth_microservice/server/usecase"
	"PLANEXA_backend/models"
)

type SessionUseCaseImpl struct {
	sessRepo repository.SessionRedis
}

func CreateSessionUseCase(sessionRepository repository.SessionRedis) usecase.SessionUseCase {
	return &SessionUseCaseImpl{sessRepo: sessionRepository}
}

func (sessUseCase *SessionUseCaseImpl) SetSession(session models.Session) error {
	return sessUseCase.sessRepo.SetSession(session)
}

func (sessUseCase *SessionUseCaseImpl) GetSession(cookieVal string) (uint64, error) {
	return sessUseCase.sessRepo.GetSession(cookieVal)
}

func (sessUseCase *SessionUseCaseImpl) DeleteSession(cookieVal string) error {
	return sessUseCase.sessRepo.DeleteSession(cookieVal)
}
