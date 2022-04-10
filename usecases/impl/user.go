package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/hash"
	"PLANEXA_backend/models"
	"PLANEXA_backend/redis"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
	"PLANEXA_backend/utils"
	"github.com/microcosm-cc/bluemonday"
)

type UserUseCaseImpl struct {
	rep *repositories.UserRepository
	red *planexa_redis.RedisConnect
}

func MakeUserUsecase(rep_ *repositories.UserRepository, red_ *planexa_redis.RedisConnect) usecases.UserUseCase {
	return &UserUseCaseImpl{rep: rep_, red: red_}
}

func (userUseCase *UserUseCaseImpl) Login(user models.User) (uint, string, error) {
	// вызываю из бд проверку есть ли юзер
	//сравниваю пароли
	isAble, err := userUseCase.rep.IsAbleToLogin(user.Username, user.Password)
	if err != nil {
		return 0, "", err
	}
	newUser, err := userUseCase.rep.GetUserByLogin(user.Username)
	if err != nil {
		return 0, "", err
	}

	if !isAble {
		return 0, "", customErrors.ErrUnauthorized
	}

	token := utils.GenerateSessionToken()
	err = userUseCase.red.SetSession(models.Session{UserId: newUser.IdU, CookieValue: token})
	// сохраняю сессию в бд и возвращаю token
	if err != nil {
		return 0, "", err
	}
	return newUser.IdU, token, nil
}

func (userUseCase *UserUseCaseImpl) Register(user models.User) (uint, string, error) {
	err := utils.CheckPassword(user.Password)
	if err != nil {
		return 0, "", err
	}

	// проверяю в БД есть ли такой юзер и обрабатываю ошибку в случае чего

	isExist, err := userUseCase.rep.IsExist(user.Username)
	if isExist {
		return 0, "", customErrors.ErrUsernameExist
	} else if err != nil {
		return 0, "", err
	}

	hashPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return 0, "", err
	}
	//задаем текущему пользователю "новый" пароль
	user.Password = hashPassword

	// добавляю юзера в бд и создаю токен для него, добавляю в бд сессию

	userId, err := userUseCase.rep.Create(&user)
	if err != nil {
		return 0, "", err
	}

	token := utils.GenerateSessionToken()
	err = userUseCase.red.SetSession(models.Session{UserId: user.IdU, CookieValue: token})
	// возвращаю токен и ошибку
	return userId, token, err
}

func (userUseCase *UserUseCaseImpl) Logout(token string) error {
	return userUseCase.red.DeleteSession(token)
}

func (userUseCase *UserUseCaseImpl) GetInfo(userId uint) (models.User, error) {
	// получаю из бд всю инфу по айдишнику кроме пароля
	user, err := userUseCase.rep.GetUserById(userId)
	sanitizer := bluemonday.UGCPolicy()
	user.Password = sanitizer.Sanitize(user.Password)
	user.Username = sanitizer.Sanitize(user.Username)
	user.ImgAvatar = sanitizer.Sanitize(user.ImgAvatar)

	user.Password = ""
	return *user, err
}
