package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/user_microservice/server/handler"
	"PLANEXA_backend/user_microservice/server/usecase"
	"context"
	"encoding/json"
)

type UserServerImpl struct {
	userUseCase usecase.UserUseCase
	handler.UnimplementedUserServiceServer
}

func CreateSessionServer(userUseCase usecase.UserUseCase) handler.UserServiceServer {
	return &UserServerImpl{userUseCase: userUseCase}
}

func (userServ *UserServerImpl) Create(ctx context.Context, in *handler.User) (*handler.IdUser, error) {
	if in == nil {
		return &handler.IdUser{}, customErrors.ErrBadInputData
	}
	var boards []models.Board
	err := json.Unmarshal(in.BOARDS, &boards)
	if err != nil {
		return &handler.IdUser{}, err
	}
	userId, err := userServ.userUseCase.Create(&models.User{Username: in.UserData.Uname.USERNAME,
		Password: in.UserData.Pass, IdU: uint(in.IDU.IDU), ImgAvatar: in.IMG, Boards: boards})
	return &handler.IdUser{IDU: uint64(userId)}, err
}

func (userServ *UserServerImpl) Update(ctx context.Context, in *handler.User) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}
	var boards []models.Board
	err := json.Unmarshal(in.BOARDS, &boards)
	if err != nil {
		return &handler.Nothing{}, err
	}
	err = userServ.userUseCase.Update(&models.User{Username: in.UserData.Uname.USERNAME,
		Password: in.UserData.Pass, IdU: uint(in.IDU.IDU), ImgAvatar: in.IMG, Boards: boards})
	return &handler.Nothing{Dummy: true}, err
}

func (userServ *UserServerImpl) IsAbleToLogin(ctx context.Context, in *handler.CheckLog) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}

	is, err := userServ.userUseCase.IsAbleToLogin(in.Pass, in.Uname.USERNAME)
	return &handler.Nothing{Dummy: is}, err
}

func (userServ *UserServerImpl) AddUserToBoard(ctx context.Context, in *handler.Ids) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}

	err := userServ.userUseCase.AddUserToBoard(uint(in.IDU.IDU), uint(in.IDB.IDB))
	return &handler.Nothing{}, err
}

func (userServ *UserServerImpl) GetUserByLogin(ctx context.Context, in *handler.Username) (*handler.User, error) {
	if in == nil {
		return &handler.User{}, customErrors.ErrBadInputData
	}

	user, err := userServ.userUseCase.GetUserByLogin(in.USERNAME)
	if err != nil {
		return &handler.User{}, err
	}
	boardsBytes, err := json.Marshal(user.Boards)
	return &handler.User{IDU: &handler.IdUser{IDU: uint64(user.IdU)},
		UserData: &handler.CheckLog{Uname: &handler.Username{USERNAME: user.Username}, Pass: user.Password},
		IMG:      user.ImgAvatar, BOARDS: boardsBytes}, err
}

func (userServ *UserServerImpl) GetUserById(ctx context.Context, in *handler.IdUser) (*handler.User, error) {
	if in == nil {
		return &handler.User{}, customErrors.ErrBadInputData
	}

	user, err := userServ.userUseCase.GetUserById(uint(in.IDU))
	if err != nil {
		return &handler.User{}, err
	}
	boardsBytes, err := json.Marshal(user.Boards)
	return &handler.User{IDU: &handler.IdUser{IDU: uint64(user.IdU)},
		UserData: &handler.CheckLog{Uname: &handler.Username{USERNAME: user.Username}, Pass: user.Password},
		IMG:      user.ImgAvatar, BOARDS: boardsBytes}, err
}

func (userServ *UserServerImpl) IsExist(ctx context.Context, in *handler.Username) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, customErrors.ErrBadInputData
	}

	is, err := userServ.userUseCase.IsExist(in.USERNAME)
	return &handler.Nothing{Dummy: is}, err
}

func (userServ *UserServerImpl) GetUsersLike(ctx context.Context, in *handler.Username) (*handler.Users, error) {
	if in == nil {
		return &handler.Users{}, customErrors.ErrBadInputData
	}
	users, err := userServ.userUseCase.GetUsersLike(in.USERNAME)
	if err != nil {
		return &handler.Users{}, err
	}
	bytesUsers, err := json.Marshal(users)
	return &handler.Users{USERS: bytesUsers}, err
}
