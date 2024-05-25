package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	pb "github.com/emzola/ridewise/gatewayservice/genproto"
	"github.com/emzola/ridewise/gatewayservice/internal/controller/gateway"
	authenticationGateway "github.com/emzola/ridewise/gatewayservice/internal/gateway/authenticationservice/grpc"
	smsNotificationGateway "github.com/emzola/ridewise/gatewayservice/internal/gateway/smsnotificationservice/grpc"
	grpcHandler "github.com/emzola/ridewise/gatewayservice/internal/handler/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	baseConfig := filepath.Join("gatewayservice", "configs", "base.yaml")
	f, err := os.Open(baseConfig)
	if err != nil {
		logger.Error("failed to open configuration file", slog.Any("error", err))
	}
	defer f.Close()

	var cfg config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		logger.Error("failed to parse configuration", slog.Any("error", err))
	}
	port := cfg.API.Port
	logger.Info("starting the gateway service", slog.Int("port", port))

	_, cancel := context.WithCancel(context.Background())

	ctrl := gateway.New(
		authenticationGateway.New(),
		smsNotificationGateway.New(),
	)
	handler := grpcHandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		logger.Error("failed to listen", slog.Any("error", err))
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	pb.RegisterGatewayServiceServer(srv, handler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := <-stop
		cancel()
		logger.Info("shutting down gracefully", slog.String("signal", s.String()))
		srv.GracefulStop()
		logger.Info("server stopped")
	}()

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

	wg.Wait()
}
