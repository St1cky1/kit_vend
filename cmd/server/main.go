package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/St1cky1/kit_vend/internal/api/kit_vending"
	grpcserver "github.com/St1cky1/kit_vend/internal/grpc"
	"github.com/St1cky1/kit_vend/internal/storage"
	"github.com/St1cky1/kit_vend/internal/usecase"
	pbv1 "github.com/St1cky1/kit_vend/pb/v1"
	"github.com/St1cky1/kit_vend/pkg/config"
)

func main() {
	cfg := config.Load()
	log := grpcserver.NewLogger(cfg.LogLevel)

	log.Info("Starting Kit Vending Backend", "http_port", cfg.Server.Port)

	kitClient := kit_vending.NewClient(cfg.KitVendingAPI.CompanyId, cfg.KitVendingAPI.Login, cfg.KitVendingAPI.Password)
	if cfg.LogLevel == "debug" {
		kitClient.SetDebug(true)
		log.Info("Debug mode enabled - Kit Vending API responses will be logged")
	}

	vmRepo := storage.NewMockVendingMachineRepository()
	saleRepo := storage.NewMockSaleRepository()
	actionRepo := storage.NewMockActionRepository()
	eventRepo := storage.NewMockEventRepository()
	vmStateRepo := storage.NewMockVMStateRepository()
	remainsRepo := storage.NewMockVendingMachineRemainsRepository()

	uc := usecase.NewVendingMachineUseCase(kitClient, vmRepo, saleRepo, actionRepo, eventRepo, vmStateRepo, remainsRepo)

	grpcPort := ":50051"
	httpPort := ":" + cfg.Server.Port

	grpcListener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error("failed to listen on gRPC port", "error", err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcserver.UnaryServerLoggingInterceptor(log)),
		grpc.StreamInterceptor(grpcserver.StreamServerLoggingInterceptor(log)),
	)

	vmService := grpcserver.NewVendingMachineServiceServer(uc)
	pbv1.RegisterVendingMachineServiceServer(grpcServer, vmService)

	go func() {
		log.Info("gRPC server listening", "address", "0.0.0.0:50051")
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Error("gRPC server error", "error", err)
		}
	}()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error("failed to dial gRPC server", "error", err)
		return
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()
	if err := pbv1.RegisterVendingMachineServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Error("failed to register gateway", "error", err)
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", gwmux)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
			log.Error("failed to write health response", "error", err)
		}
	})

	log.Info("HTTP server started", "address", "0.0.0.0"+httpPort)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Error("HTTP server error", "error", err)
	}
}
