package main

import (
	"PLANEXA_backend/metrics"
	"PLANEXA_backend/user_microservice/server_user/handler"
	service_impl "PLANEXA_backend/user_microservice/server_user/handler/impl"
	repository_impl "PLANEXA_backend/user_microservice/server_user/repository/impl"
	usecase_impl "PLANEXA_backend/user_microservice/server_user/usecase/impl"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"net/http"
	"time"
)

func Run() {
	userRepo := repository_impl.CreateUserRep()
	userUseCase := usecase_impl.CreateUserUseCase(userRepo)
	listener, err := net.Listen("tcp", "2022_1_samoekrnaz_user_microservice_1:8083")
	if err != nil {
		return
	}

	grpcSrv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	handler.RegisterUserServiceServer(grpcSrv, service_impl.CreateUserServer(userUseCase))

	prometheus.MustRegister(metrics.User)
	prometheus.MustRegister(metrics.DurationUser)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":8084", nil))
	}()

	if err = grpcSrv.Serve(listener); err != nil {
		return
	}
}
