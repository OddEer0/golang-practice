package main

import (
	pgxPractice "github.com/OddEer0/golang-practice/practices/pgx_practice"
	pgxHandler "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_handler"
	pgxInteractor "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_interactor"
	proto "github.com/OddEer0/golang-practice/protogen"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func main() {
	cfg := pgxPractice.NewConfig()
	lis, err := net.Listen("tcp", cfg.Server.Address)
	connPool := pgxPractice.ConnectPg(cfg)
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			slog.Error("close tcp connection error", slog.String("cause", err.Error()))
		}
	}(lis)
	if err != nil {
		slog.Error("net listen tcp error", slog.String("cause", err.Error()))
		return
	}

	grpcServer := grpc.NewServer()
	interactor := pgxInteractor.New(connPool)
	proto.RegisterNewsServiceServer(grpcServer, pgxHandler.NewGrpcHandler(interactor))
	ch := make(chan struct{})
	go func() {
		slog.Info("grpc server start", slog.String("address", cfg.Server.Address))
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("grpc serve error", slog.String("cause", err.Error()))
			ch <- struct{}{}
		}
	}()
	<-ch
}
