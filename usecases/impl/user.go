package impl

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/redis"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/utils"
)

type UserUsecase struct {
	rep *repositories.UserRepository
	red *planexa_redis.RedisConnect
}

func MakeUsecase(rep_ *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{rep: rep_}
}

func (userUsecase *UserUsecase) Login(user models.User) (string, error) {
	// вызываю из бд проверку есть ли юзер
	//сравниваю пароли

	// сохраняю сессию в бд и возвращаю token

	return "", nil
}

func Register(user models.User) (string, error) {
	err := utils.CheckPassword(user.Password)
	if err != nil {
		return "", err
	}

	// проверяю в БД есть ли такой юзер и обрабатываю ошибку в случае чего

	// добавляю юзера в бд и создаю токен для него, добавляю в бд сессию

	// возвращаю токен и ошибку
	return "token", nil
}

func Logout(token string) error {
	// достаю из бд все сессии и проверяю на наличие
	//если есть, удаляю из бд
	var err error
	return err
}

func GetInfo(userId uint) (models.User, error) {
	// получаю из бд всю инфу по айдишнику кроме пароля
	var err error
	return models.User{}, err
}
