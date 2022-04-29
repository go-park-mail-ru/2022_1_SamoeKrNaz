package main

import (
	"PLANEXA_backend/auth_microservice/server/handler"
	"PLANEXA_backend/auth_microservice/server/handler/impl"
	repository_impl "PLANEXA_backend/auth_microservice/server/repository/impl"
	usecase_impl "PLANEXA_backend/auth_microservice/server/usecase/impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

func Run() {
	redis := repository_impl.CreateSessRep()
	sessUseCase := usecase_impl.CreateSessionUseCase(redis)
	listener, err := net.Listen("tcp", "session:8081")
	if err != nil {
		return
	}

	grpcSrv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	handler.RegisterAuthCheckerServer(grpcSrv, impl.CreateSessionServer(sessUseCase))
	if err = grpcSrv.Serve(listener); err != nil {
		return
	}
}
