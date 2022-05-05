package main

import (
	"PLANEXA_backend/metrics"
	"PLANEXA_backend/user_microservice/server_user_ms/handler"
	service_impl "PLANEXA_backend/user_microservice/server_user_ms/handler/impl"
	repository_impl "PLANEXA_backend/user_microservice/server_user_ms/repository/impl"
	usecase_impl "PLANEXA_backend/user_microservice/server_user_ms/usecase/impl"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"time"
)

func Run() {
	newDb, err := gorm.Open(postgres.Open("host=postgres user=Planexa password=WEB21Planexa dbname=DB_Planexa port=5432"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userRepo := repository_impl.CreateUserRep(newDb)
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
