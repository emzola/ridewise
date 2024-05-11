package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"

	"github.com/emzola/ridewise/gen"
	"github.com/emzola/ridewise/riderservice/internal/controller/rider"
	grpcHandler "github.com/emzola/ridewise/riderservice/internal/handler/grpc"
	"github.com/emzola/ridewise/riderservice/internal/repository/memory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	baseConfig := filepath.Join("riderservice", "configs", "base.yaml")
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
	logger.Info("starting the rider service", slog.Int("port", port))

	repo := memory.New()
	ctrl := rider.New(repo)
	handler := grpcHandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		logger.Error("failed to listen", slog.Any("error", err))
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterRiderServiceServer(srv, handler)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
