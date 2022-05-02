package impl

import (
	"PLANEXA_backend/errors"
	"PLANEXA_backend/hash"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/user_microservice/server/handler"
	"context"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"gorm.io/gorm"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

const filePathAvatars = "avatars/"

type UserRepositoryImpl struct {
	db     *gorm.DB
	client handler.UserServiceClient
	ctx    context.Context
}

func MakeUserRepository(db *gorm.DB, cl handler.UserServiceClient) repositories.UserRepository {
	return &UserRepositoryImpl{db: db, client: cl, ctx: context.Background()}
}

func (userRepository *UserRepositoryImpl) Create(user *models.User) (uint, error) {
	// проверка на уже существующего пользователя
	user.ImgAvatar = filePathAvatars + "default.webp"
	err := userRepository.db.Create(user).Error
	return user.IdU, err
}

func (userRepository *UserRepositoryImpl) Update(user *models.User) error {
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
		} else if err != nil && err != customErrors.ErrUserNotFound {
			return err
		} else {
			currentData.Username = user.Username
		}
	}
	// если мы поменяли пароль, то надо его захешировать
	if !hash.CheckPasswordHash(user.Password, currentData.Password) && user.Password != "" {
		currentData.Password, err = hash.HashPassword(user.Password)
		if err != nil {
			return err
		}
	}
	return userRepository.db.Save(currentData).Error
}

func (userRepository *UserRepositoryImpl) SaveAvatar(user *models.User, header *multipart.FileHeader) error {
	if user.ImgAvatar != "" {
		currentData, err := userRepository.GetUserById(user.IdU)
		if err != nil {
			return err
		}

		fileName := strings.Join([]string{filePathAvatars, strconv.Itoa(int(currentData.IdU)), ".webp"}, "")
		output, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer output.Close()

		openFile, err := header.Open()
		if err != nil {
			return err
		}

		img, _, err := image.Decode(openFile)
		if err != nil {
			return err
		}

		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
		if err != nil {
			return err
		}

		err = webp.Encode(output, img, options)
		if err != nil {
			return err
		}

		currentData.ImgAvatar = fileName
		return userRepository.db.Save(currentData).Error
	}
	return nil
}

func (userRepository *UserRepositoryImpl) IsAbleToLogin(username string, password string) (bool, error) {
	// проверка на существование пользователя по никнейму
	isExist, err := userRepository.IsExist(username)
	if !isExist {
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

func (userRepository *UserRepositoryImpl) AddUserToBoard(IdB uint, IdU uint) error {
	user, err := userRepository.GetUserById(IdU)
	if err != nil {
		return err
	}
	return userRepository.db.Model(&models.Board{IdB: IdB}).Association("Users").Append(user)
}

func (userRepository *UserRepositoryImpl) GetUserByLogin(username string) (*models.User, error) {
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

func (userRepository *UserRepositoryImpl) GetUserById(IdU uint) (*models.User, error) {
	// указатель на структуру, которую вернем
	user := new(models.User)
	result := userRepository.db.Find(user, IdU)
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

func (userRepository *UserRepositoryImpl) IsExist(username string) (bool, error) {
	result, err := userRepository.GetUserByLogin(username)
	if err != nil && err != customErrors.ErrUserNotFound {
		return false, err
	} else if result == nil {
		return false, nil
	}
	return true, nil
}

func (userRepository *UserRepositoryImpl) GetUsersLike(username string) (*[]models.User, error) {
	users := new([]models.User)
	err := userRepository.db.Where("lower(username) LIKE lower(?)", username).Limit(15).Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
