package main

import (
	"PLANEXA_backend/auth_microservice/server_session_ms/handler"
	"PLANEXA_backend/auth_microservice/server_session_ms/handler/impl"
	repository_impl "PLANEXA_backend/auth_microservice/server_session_ms/repository/impl"
	usecase_impl "PLANEXA_backend/auth_microservice/server_session_ms/usecase/impl"
	"PLANEXA_backend/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"net/http"
	"time"
)

type Config struct {
	sessionContainer string
	metricsPort      string
	redisContainer   string
	metricsPath      string
}

func ParseConfig() (conf Config) {
	viper.AddConfigPath("./auth_microservice/server_session_ms/")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	conf.sessionContainer = viper.GetString("sessionContainer")
	conf.redisContainer = viper.GetString("redisContainer")
	conf.metricsPort = viper.GetString("metricsPort")
	conf.metricsPath = viper.GetString("metricsPath")
	return
}

func Run() {
	conf := ParseConfig()
	redis := repository_impl.CreateSessRep(conf.redisContainer)
	sessUseCase := usecase_impl.CreateSessionUseCase(redis)
	listener, err := net.Listen("tcp", conf.sessionContainer)
	if err != nil {
		return
	}

	grpcSrv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	handler.RegisterAuthCheckerServer(grpcSrv, impl.CreateSessionServer(sessUseCase))

	prometheus.MustRegister(metrics.Session)
	prometheus.MustRegister(metrics.DurationSession)

	http.Handle(conf.metricsPath, promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(conf.metricsPort, nil))
	}()

	if err = grpcSrv.Serve(listener); err != nil {
		return
	}
}
