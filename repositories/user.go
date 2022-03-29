package repositories

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/hash"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (userRepository *UserRepository) Create(user *models.User) error {
	// проверка на уже существующего пользователя
	isExist := userRepository.IsExist(user.Username)
	if isExist != nil {
		return isExist
	}
	// пароли нужно хранить скрытно, поэтому хешируем
	hashPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}
	//задаем текущему пользователю "новый" пароль
	user.Password = hashPassword

	err = userRepository.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) Update(user *models.User) (err error) {
	// будем предполагать, что пришла структура с новыми полями, и мог измениться никнейм
	// поэтому поиск по никнейму ничего не даст, будем искать по Id
	currentData := userRepository.GetUserById(user.IdP)
	if currentData == nil {
		return customErrors.ErrIdNotExist
	}
	// теперь будем искать, какое поле поменялось
	if currentData.Username != user.Username {
		//проверяем, не занят ли новый никнейм
		isExist := userRepository.IsExist(user.Username)
		//если такой никнейм уже занят, то отправляем ошибку
		if isExist != nil {
			return isExist
		}
	}
	// если мы поменяли пароль, то надо его захешировать
	if hash.CheckPasswordHash(user.Password, currentData.Password) {
		user.Password, err = hash.HashPassword(user.Password)
		if err != nil {
			return err
		}
	}
	userRepository.db.Save(user)
	return nil
}

func (userRepository *UserRepository) Login(username string, password string) (*models.User, error) {
	// проверка на существование пользователя по никнейму
	isExist := userRepository.IsExist(username)
	if isExist != nil {
		return nil, isExist
	}
	// вычисляем хеш от пароля
	hashPassword, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := new(models.User)
	// чекаем в базе правильность данных
	result := userRepository.db.Select("*").Where("username = ?", username, "password = ?", hashPassword).Find(&user)
	// если выборка в 0 строк, то не сошлись данные
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrBadInputData
	} else {
		return user, nil
	}
}

func (userRepository *UserRepository) GetUserByLogin(username string) *models.User {
	// указатель на структуру, которую вернем
	user := new(models.User)
	result := userRepository.db.Select("*").Where("username = ?", username).Find(&user)
	// если выборка в 0 строк, то такого пользователя нет
	if result.RowsAffected == 0 {
		return nil
	} else {
		// иначе вернем пользователя
		return user
	}
}

func (userRepository *UserRepository) GetUserById(IdP uint) *models.User {
	// указатель на структуру, которую вернем
	user := new(models.User)
	result := userRepository.db.Select("*").Where("IdP = ?", IdP).Find(&user)
	// если выборка в 0 строк, то такого пользователя нет
	if result.RowsAffected == 0 {
		return nil
	} else {
		// иначе вернем пользователя
		return user
	}
}

func (userRepository *UserRepository) IsExist(username string) error {
	// пробуем найти такого пользователя
	result := userRepository.db.Select("*").Where("username = ?", username)
	// если выборка в 0 строк, то ошибочка
	if result.RowsAffected == 0 {
		return customErrors.ErrUsernameExist
	} else {
		return nil
	}
}
