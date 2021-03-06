package impl

import (
	"PLANEXA_backend/errors"
	"PLANEXA_backend/hash"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"PLANEXA_backend/utils"
	"github.com/microcosm-cc/bluemonday"
	"mime/multipart"
	"strconv"
	"strings"
)

type UserUseCaseImpl struct {
	rep repositories.UserRepository
	red repositories.SessionRepository
}

func MakeUserUsecase(rep_ repositories.UserRepository, red_ repositories.SessionRepository) usecases.UserUseCase {
	return &UserUseCaseImpl{rep: rep_, red: red_}
}

func (userUseCase *UserUseCaseImpl) Login(user models.User) (uint, string, error) {
	// вызываю из бд проверку есть ли юзер
	//сравниваю пароли
	isAble, err := userUseCase.rep.IsAbleToLogin(user.Username, user.Password)
	if err != nil {
		return 0, "", err
	}
	if !isAble {
		return 0, "", customErrors.ErrUnauthorized
	}

	newUser, err := userUseCase.rep.GetUserByLogin(user.Username)
	if err != nil {
		return 0, "", err
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
	err = userUseCase.red.SetSession(models.Session{UserId: userId, CookieValue: token})
	// возвращаю токен и ошибку
	return userId, token, err
}

func (userUseCase *UserUseCaseImpl) Logout(token string) error {
	err := userUseCase.red.DeleteSession(token)
	return err
}

func (userUseCase *UserUseCaseImpl) GetInfoById(userId uint) (models.User, error) {
	// получаю из бд всю инфу по айдишнику кроме пароля
	user, err := userUseCase.rep.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	sanitizer := bluemonday.UGCPolicy()
	user.Password = sanitizer.Sanitize(user.Password)
	user.Username = sanitizer.Sanitize(user.Username)
	user.ImgAvatar = sanitizer.Sanitize(user.ImgAvatar)

	user.Password = ""
	return *user, err
}

func (userUseCase *UserUseCaseImpl) SaveAvatar(user *models.User, header *multipart.FileHeader) (string, error) {
	err := userUseCase.rep.SaveAvatar(user, header)
	return strings.Join([]string{strconv.Itoa(int(user.IdU)), ".webp"}, ""), err
}

func (userUseCase *UserUseCaseImpl) RefactorProfile(user models.User) error {
	return userUseCase.rep.Update(&user)
}

func (userUseCase *UserUseCaseImpl) GetUsersLike(username string) (*[]models.User, error) {
	return userUseCase.rep.GetUsersLike(username)
}
