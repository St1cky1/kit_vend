package main

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"

	apiv1 "github.com/St1cky1/kit_vend/api/v1"
	"github.com/St1cky1/kit_vend/internal/api"
	grpcserver "github.com/St1cky1/kit_vend/internal/grpc"
	"github.com/St1cky1/kit_vend/internal/handler"
	"github.com/St1cky1/kit_vend/internal/storage"
	"github.com/St1cky1/kit_vend/internal/usecase"
	"github.com/St1cky1/kit_vend/pkg/config"
	"github.com/St1cky1/kit_vend/pkg/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	log.Info("Starting Kit Vending Backend", "port", cfg.Server.Port)

	kitClient := api.NewClient(cfg.KitVendingAPI.CompanyId, cfg.KitVendingAPI.Login, cfg.KitVendingAPI.Password)
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

	go startGRPCServer(uc, log)

	startHTTPServer(uc, log, cfg.Server.Port)
}

func startGRPCServer(uc *usecase.VendingMachineUseCase, log *logger.Logger) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error("failed to listen on gRPC port", "error", err)
		return
	}

	grpcServer := grpc.NewServer()
	vmService := grpcserver.NewVendingMachineServiceServer(uc)
	apiv1.RegisterVendingMachineServiceServer(grpcServer, vmService)

	log.Info("gRPC server listening", "address", "0.0.0.0:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Error("gRPC server error", "error", err)
	}
}

func startHTTPServer(uc *usecase.VendingMachineUseCase, log *logger.Logger, port string) {
	h := handler.NewVendingMachineHandler(uc, log)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
			log.Error("failed to write health response", "error", err)
		}
	})

	r.Get("/api/v1/vending-machines/{id}", h.GetVendingMachineByID)
	r.Get("/api/v1/sales", h.GetSales)
	r.Get("/api/v1/actions", h.GetActions)
	r.Get("/api/v1/vm-states", h.GetVMStates)
	r.Get("/api/v1/events", h.GetEvents)
	r.Post("/api/v1/commands", h.SendCommand)
	r.Get("/api/v1/vending-machines/{id}/remains", h.GetVendingMachineRemains)

	log.Info("HTTP server started", "address", "0.0.0.0:"+port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Error("HTTP server error", "error", err)
	}
}
