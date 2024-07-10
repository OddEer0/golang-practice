package main

import (
	"context"
	pgxPractice "github.com/OddEer0/golang-practice/practices/pgx_practice"
	pgxHandler "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_handler"
	pgxInteractor "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_interactor"
	proto "github.com/OddEer0/golang-practice/protogen"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
)

func main() {
	slogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(slogger)
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

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		err := proto.RegisterNewsServiceHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
		if err != nil {
			panic(err)
		}
		log.Printf("server listening at 5501")
		if err := http.ListenAndServe(":5501", mux); err != nil {
			panic(err)
		}
	}()

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
