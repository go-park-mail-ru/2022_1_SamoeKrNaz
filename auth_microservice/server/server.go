package main

import (
	"PLANEXA_backend/auth_microservice/server/handler"
	"PLANEXA_backend/auth_microservice/server/handler/impl"
	repository_impl "PLANEXA_backend/auth_microservice/server/repository/impl"
	usecase_impl "PLANEXA_backend/auth_microservice/server/usecase/impl"
	"PLANEXA_backend/metrics"
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
	redis := repository_impl.CreateSessRep()
	sessUseCase := usecase_impl.CreateSessionUseCase(redis)
	listener, err := net.Listen("tcp", "2022_1_samoekrnaz_session_1:8081")
	if err != nil {
		return
	}

	grpcSrv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	handler.RegisterAuthCheckerServer(grpcSrv, impl.CreateSessionServer(sessUseCase))

	prometheus.MustRegister(metrics.Session)
	prometheus.MustRegister(metrics.Duration)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":8082", nil))
	}()

	if err = grpcSrv.Serve(listener); err != nil {
		return
	}
}
