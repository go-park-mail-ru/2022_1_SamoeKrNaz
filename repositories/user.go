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

func (userRepository *UserRepository) MakeRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) Create(user *models.User) error {
	// проверка на уже существующего пользователя
	isExist, err := userRepository.IsExist(user.Username)
	if isExist {
		return customErrors.ErrUsernameExist
	} else if err != nil {
		return err
	}
	// пароли нужно хранить скрытно, поэтому хешируем
	hashPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}
	//задаем текущему пользователю "новый" пароль
	user.Password = hashPassword
	return userRepository.db.Create(user).Error
}

func (userRepository *UserRepository) Update(user *models.User) error {
	// будем предполагать, что пришла структура с новыми полями, и мог измениться никнейм
	// поэтому поиск по никнейму ничего не даст, будем искать по Id
	currentData, err := userRepository.GetUserById(user.IdU)
	if err != nil {
		return err
	}
	// теперь будем искать, какое поле поменялось
	if currentData.Username != user.Username && user.Username != "" {
		//проверяем, не занят ли новый никнейм
		isExist, err := userRepository.IsExist(user.Username)
		//если такой никнейм уже занят, то отправляем ошибку
		if isExist {
			return customErrors.ErrUsernameExist
		} else if err != nil {
			return err
		} else {
			currentData.Username = user.Username
		}
	}
	// если мы поменяли пароль, то надо его захешировать
	if hash.CheckPasswordHash(user.Password, currentData.Password) && user.Password != "" {
		currentData.Password, err = hash.HashPassword(user.Password)
		if err != nil {
			return err
		}
	}
	return userRepository.db.Save(currentData).Error
}

func (userRepository *UserRepository) IsAbleToLogin(username string, password string) (bool, error) {
	// проверка на существование пользователя по никнейму
	isExist, err := userRepository.IsExist(username)
	if isExist != true {
		return false, customErrors.ErrUsernameNotExist
	}
	if err != nil {
		return false, err
	}
	// чекаем в базе правильность данных
	user, err := userRepository.GetUserByLogin(username)
	if err != nil {
		return false, err
	}
	// если выборка в 0 строк, то не сошлись данные
	if user == nil {
		return false, customErrors.ErrBadInputData
	} else if hash.CheckPasswordHash(password, user.Password) {
		// проверим правильность пароля
		return true, nil
	} else {
		return false, customErrors.ErrBadInputData
	}
}

func (userRepository *UserRepository) AddUserToBoard(IdB uint, IdU uint) error {
	user, err := userRepository.GetUserById(IdU)
	if err != nil {
		return err
	}
	return userRepository.db.Model(&models.Board{IdB: IdB}).Association("Users").Append(user)
}

func (userRepository *UserRepository) GetUserByLogin(username string) (*models.User, error) {
	// указатель на структуру, которую вернем
	user := new(models.User)
	result := userRepository.db.Where("username = ?", username).Find(user)
	// если выборка в 0 строк, то такого пользователя нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrUserNotFound
	} else {
		// иначе вернем пользователя
		return user, nil
	}
}

func (userRepository *UserRepository) GetUserById(IdP uint) (*models.User, error) {
	// указатель на структуру, которую вернем
	user := new(models.User)
	result := userRepository.db.Find(user, IdP)
	// если выборка в 0 строк, то такого пользователя нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrUserNotFound
	} else if result.Error != nil {
		// если произошла ошибка при выборке
		return nil, result.Error
	} else {
		// иначе вернем пользователя
		return user, nil
	}
}

func (userRepository *UserRepository) IsExist(username string) (bool, error) {
	result, err := userRepository.GetUserByLogin(username)
	if err != nil {
		return false, err
	} else if result == nil {
		return false, nil
	}
	return true, nil
}