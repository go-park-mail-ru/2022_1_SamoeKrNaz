package usecases

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/utils"
)

func Login(user models.User) (string, error) {
	// вызываю из бд проверку есть ли юзер
	var err error // ошибка из бд
	//сравниваю пароли

	// сохраняю сессию в бд и возвращаю token

	return "", err
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
