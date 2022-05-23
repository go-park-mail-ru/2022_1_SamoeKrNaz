package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"PLANEXA_backend/user_microservice/server_user_ms/handler"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
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
	boardBytes, err := json.Marshal(user)
	if err != nil {
		return 0, err
	}
	idU, err := userRepository.client.Create(userRepository.ctx, &handler.User{IDU: &handler.IdUser{IDU: uint64(user.IdU)},
		UserData: &handler.CheckLog{Pass: user.Password, Uname: &handler.Username{USERNAME: user.Username}}, IMG: user.ImgAvatar,
		BOARDS: boardBytes})
	return uint(idU.IDU), err
}

func (userRepository *UserRepositoryImpl) Update(user *models.User) error {
	boardBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = userRepository.client.Create(userRepository.ctx, &handler.User{IDU: &handler.IdUser{IDU: uint64(user.IdU)},
		UserData: &handler.CheckLog{Pass: user.Password, Uname: &handler.Username{USERNAME: user.Username}}, IMG: user.ImgAvatar,
		BOARDS: boardBytes})
	return err
}

func (userRepository *UserRepositoryImpl) SaveAvatar(user *models.User, header *multipart.FileHeader) error {
	if user.ImgAvatar != "" {
		//currentData, err := userRepository.GetUserById(user.IdU)
		//if err != nil {
		//	return err
		//}
		//
		//fileName := strings.Join([]string{filePathAvatars, strconv.Itoa(int(currentData.IdU)), ".webp"}, "")
		//output, err := os.Create(fileName)
		//if err != nil {
		//	return err
		//}
		//defer output.Close()
		//
		//openFile, err := header.Open()
		//if err != nil {
		//	return err
		//}
		//
		//img, _, err := image.Decode(openFile)
		//if err != nil {
		//	return err
		//}
		//
		//options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
		//if err != nil {
		//	return err
		//}
		//
		//err = webp.Encode(output, img, options)
		//if err != nil {
		//	return err
		//}
		//
		//currentData.ImgAvatar = fileName
		//return userRepository.db.Save(currentData).Error
	}
	return nil
}

func (userRepository *UserRepositoryImpl) IsAbleToLogin(username string, password string) (bool, error) {
	isAble, err := userRepository.client.IsAbleToLogin(userRepository.ctx, &handler.CheckLog{Pass: password, Uname: &handler.Username{USERNAME: username}})
	if err != nil {
		if strings.Contains(err.Error(), customErrors.ErrUsernameNotExist.Error()) {
			err = customErrors.ErrBadInputData
		}
		if strings.Contains(err.Error(), customErrors.ErrBadInputData.Error()) {
			err = customErrors.ErrBadInputData
		}
		return false, err
	}
	return isAble.Dummy, err
}

func (userRepository *UserRepositoryImpl) AddUserToBoard(IdB uint, IdU uint) error {
	_, err := userRepository.client.AddUserToBoard(userRepository.ctx, &handler.Ids{IDU: &handler.IdUser{IDU: uint64(IdU)},
		IDB: &handler.IdBoard{IDB: uint64(IdB)}})
	return err
}

func (userRepository *UserRepositoryImpl) GetUserByLogin(username string) (*models.User, error) {
	// указатель на структуру, которую вернем

	user, err := userRepository.client.GetUserByLogin(userRepository.ctx, &handler.Username{USERNAME: username})
	if err != nil {
		return nil, err
	}
	var boards []models.Board
	err = json.Unmarshal(user.BOARDS, &boards)
	return &models.User{Username: user.UserData.Uname.USERNAME, Password: user.UserData.Pass, IdU: uint(user.IDU.IDU),
		ImgAvatar: user.IMG}, err
}

func (userRepository *UserRepositoryImpl) GetUserById(IdU uint) (*models.User, error) {
	user, err := userRepository.client.GetUserById(userRepository.ctx, &handler.IdUser{IDU: uint64(IdU)})
	if err != nil {
		return nil, err
	}
	var boards []models.Board
	err = json.Unmarshal(user.BOARDS, &boards)
	return &models.User{Username: user.UserData.Uname.USERNAME, Password: user.UserData.Pass, IdU: uint(user.IDU.IDU),
		ImgAvatar: user.IMG}, err
}

func (userRepository *UserRepositoryImpl) IsExist(username string) (bool, error) {
	isEx, err := userRepository.client.IsExist(userRepository.ctx, &handler.Username{USERNAME: username})
	return isEx.Dummy, err
}

func (userRepository *UserRepositoryImpl) GetUsersLike(username string) (*[]models.User, error) {
	usersByte, err := userRepository.client.GetUsersLike(userRepository.ctx, &handler.Username{USERNAME: username})
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = json.Unmarshal(usersByte.USERS, &users)
	return &users, err
}
