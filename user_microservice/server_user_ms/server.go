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
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	postgresHost     string
	postgresUser     string
	postgresPassword string
	postgresDbName   string
	postgresPort     string

	userContainer string

	metricsPath string
	metricsPort string
}

func ParseConfig() (conf Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	conf.postgresHost = viper.GetString("postgresHost")
	conf.postgresUser = viper.GetString("postgresUser")
	conf.postgresPassword = viper.GetString("postgresPassword")
	conf.postgresDbName = viper.GetString("postgresDbName")
	conf.postgresPort = viper.GetString("postgresPort")

	conf.userContainer = viper.GetString("userContainer")

	conf.metricsPath = viper.GetString("metricsPath")
	conf.metricsPort = viper.GetString("metricsPort")
	return
}

func Run() {
	conf := ParseConfig()
	newDb, err := gorm.Open(postgres.Open(
		strings.Join([]string{"host=", conf.postgresHost, " user=", conf.postgresUser, " password=", conf.postgresPassword, " dbname=", conf.postgresDbName, " port=", conf.postgresPort}, "")))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userRepo := repository_impl.CreateUserRep(newDb)
	userUseCase := usecase_impl.CreateUserUseCase(userRepo)
	listener, err := net.Listen("tcp", conf.userContainer)
	if err != nil {
		return
	}

	grpcSrv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	handler.RegisterUserServiceServer(grpcSrv, service_impl.CreateUserServer(userUseCase))

	prometheus.MustRegister(metrics.User)
	prometheus.MustRegister(metrics.DurationUser)

	http.Handle(conf.metricsPath, promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(conf.metricsPort, nil))
	}()

	if err = grpcSrv.Serve(listener); err != nil {
		return
	}
}
