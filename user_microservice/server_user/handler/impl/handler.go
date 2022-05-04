package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/metrics"
	"PLANEXA_backend/models"
	"PLANEXA_backend/user_microservice/server_user/handler"
	"PLANEXA_backend/user_microservice/server_user/usecase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type UserServerImpl struct {
	userUseCase usecase.UserUseCase
	handler.UnimplementedUserServiceServer
}

func CreateUserServer(userUseCase usecase.UserUseCase) handler.UserServiceServer {
	return &UserServerImpl{userUseCase: userUseCase}
}

func (userServ *UserServerImpl) Create(ctx context.Context, in *handler.User) (*handler.IdUser, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("create"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" create user").Inc()
		return &handler.IdUser{}, customErrors.ErrBadInputData
	}
	userId, err := userServ.userUseCase.Create(&models.User{Username: in.UserData.Uname.USERNAME,
		Password: in.UserData.Pass, IdU: uint(in.IDU.IDU), ImgAvatar: in.IMG})
	if err != nil {
		metrics.User.WithLabelValues("500", "error in create user").Inc()
		return &handler.IdUser{}, err
	}
	metrics.User.WithLabelValues("200", "success in create user").Inc()
	timer.ObserveDuration()
	return &handler.IdUser{IDU: uint64(userId)}, nil
}

func (userServ *UserServerImpl) Update(ctx context.Context, in *handler.User) (*handler.NothingSec, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("update"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" update user").Inc()
		return &handler.NothingSec{}, customErrors.ErrBadInputData
	}
	err := userServ.userUseCase.Update(&models.User{Username: in.UserData.Uname.USERNAME,
		Password: in.UserData.Pass, IdU: uint(in.IDU.IDU), ImgAvatar: in.IMG})
	if err != nil {
		metrics.User.WithLabelValues("500", "error in update user").Inc()
		return &handler.NothingSec{}, err
	}
	metrics.User.WithLabelValues("200", "success in update user").Inc()
	timer.ObserveDuration()
	return &handler.NothingSec{Dummy: true}, nil
}

func (userServ *UserServerImpl) IsAbleToLogin(ctx context.Context, in *handler.CheckLog) (*handler.NothingSec, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("isabletologin"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" isabletologin user").Inc()
		return &handler.NothingSec{}, customErrors.ErrBadInputData
	}
	is, err := userServ.userUseCase.IsAbleToLogin(in.Pass, in.Uname.USERNAME)
	if err != nil {
		metrics.User.WithLabelValues("500", "error in isabletologin user").Inc()
		return &handler.NothingSec{}, err
	}
	metrics.User.WithLabelValues("200", "success in isabletologin user").Inc()
	timer.ObserveDuration()
	return &handler.NothingSec{Dummy: is}, nil
}

func (userServ *UserServerImpl) AddUserToBoard(ctx context.Context, in *handler.Ids) (*handler.NothingSec, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("addusertoboard"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" addusertoboard user").Inc()
		return &handler.NothingSec{}, customErrors.ErrBadInputData
	}
	err := userServ.userUseCase.AddUserToBoard(uint(in.IDU.IDU), uint(in.IDB.IDB))
	if err != nil {
		metrics.User.WithLabelValues("500", "error in addusertoboard user").Inc()
		return &handler.NothingSec{}, err
	}
	metrics.User.WithLabelValues("200", "success in addusertoboard user").Inc()
	timer.ObserveDuration()
	return &handler.NothingSec{}, nil
}

func (userServ *UserServerImpl) GetUserByLogin(ctx context.Context, in *handler.Username) (*handler.User, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("getuserbylogin"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" getuserbylogin user").Inc()
		return &handler.User{}, customErrors.ErrBadInputData
	}

	user, err := userServ.userUseCase.GetUserByLogin(in.USERNAME)
	fmt.Println("getuserbylogin", in.USERNAME, err)
	if err != nil {
		metrics.User.WithLabelValues("500", "error in getuserbylogin user").Inc()
		return &handler.User{}, err
	}
	boardsBytes, err := json.Marshal(user.Boards)
	if err != nil {
		metrics.User.WithLabelValues("500", "error unmarshal in getuserbylogin user").Inc()
		return &handler.User{}, err
	}
	metrics.User.WithLabelValues("200", "success in getuserbylogin user").Inc()
	timer.ObserveDuration()
	return &handler.User{IDU: &handler.IdUser{IDU: uint64(user.IdU)},
		UserData: &handler.CheckLog{Uname: &handler.Username{USERNAME: user.Username}, Pass: user.Password},
		IMG:      user.ImgAvatar, BOARDS: boardsBytes}, nil
}

func (userServ *UserServerImpl) GetUserById(ctx context.Context, in *handler.IdUser) (*handler.User, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("getuserbyid"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" getuserbyid user").Inc()
		return &handler.User{}, customErrors.ErrBadInputData
	}
	user, err := userServ.userUseCase.GetUserById(uint(in.IDU))
	fmt.Println("getuserbyid", in, err)
	if err != nil {
		metrics.User.WithLabelValues("500", "error in getuserbyid user").Inc()
		return &handler.User{}, err
	}
	boardsBytes, err := json.Marshal(user.Boards)
	if err != nil {
		metrics.User.WithLabelValues("500", "error unmarshal in getuserbyid user").Inc()
		return &handler.User{}, err
	}
	metrics.User.WithLabelValues("200", "success in getuserbyid user").Inc()
	timer.ObserveDuration()
	return &handler.User{IDU: &handler.IdUser{IDU: uint64(user.IdU)},
		UserData: &handler.CheckLog{Uname: &handler.Username{USERNAME: user.Username}, Pass: user.Password},
		IMG:      user.ImgAvatar, BOARDS: boardsBytes}, nil
}

func (userServ *UserServerImpl) IsExist(ctx context.Context, in *handler.Username) (*handler.NothingSec, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("isexist"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" isexist user").Inc()
		return &handler.NothingSec{}, customErrors.ErrBadInputData
	}
	is, err := userServ.userUseCase.IsExist(in.USERNAME)
	fmt.Println("isexist", in.USERNAME, err)
	if err != nil {
		metrics.User.WithLabelValues("500", "error in isexist user").Inc()
		return &handler.NothingSec{}, err
	}
	metrics.User.WithLabelValues("200", "success in isexist user").Inc()
	timer.ObserveDuration()
	return &handler.NothingSec{Dummy: is}, nil
}

func (userServ *UserServerImpl) GetUsersLike(ctx context.Context, in *handler.Username) (*handler.Users, error) {
	timer := prometheus.NewTimer(metrics.DurationUser.WithLabelValues("getuserslike"))
	if in == nil {
		metrics.User.WithLabelValues("500", "nil \"in\" getuserslike user").Inc()
		return &handler.Users{}, customErrors.ErrBadInputData
	}
	users, err := userServ.userUseCase.GetUsersLike(in.USERNAME)
	if err != nil {
		metrics.User.WithLabelValues("500", "error in getuserslike user").Inc()
		return &handler.Users{}, err
	}
	bytesUsers, err := json.Marshal(users)
	if err != nil {
		metrics.User.WithLabelValues("500", "error unmarshal in getuserslike user").Inc()
	}
	metrics.User.WithLabelValues("200", "success in getuserslike user").Inc()
	timer.ObserveDuration()
	return &handler.Users{USERS: bytesUsers}, nil
}
