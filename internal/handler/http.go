package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/St1cky1/kit_vend/internal/usecase"
	"github.com/St1cky1/kit_vend/pkg/logger"
)

type VendingMachineHandler struct {
	uc  *usecase.VendingMachineUseCase
	log *logger.Logger
}

func NewVendingMachineHandler(uc *usecase.VendingMachineUseCase, log *logger.Logger) *VendingMachineHandler {
	return &VendingMachineHandler{
		uc:  uc,
		log: log,
	}
}

func (h *VendingMachineHandler) GetVendingMachineByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.log.Error("missing id parameter")
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	vendingMachineID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error("invalid id parameter", "error", err)
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	vm, err := h.uc.GetVendingMachineByID(r.Context(), vendingMachineID)
	if err != nil {
		h.log.Error("failed to get vending machine", "id", vendingMachineID, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vm); err != nil {
		h.log.Error("failed to encode response", "error", err)
	}
}

func (h *VendingMachineHandler) GetSales(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from_date")
	toDate := r.URL.Query().Get("to_date")
	vendingMachineID := 0

	if vmID := r.URL.Query().Get("vending_machine_id"); vmID != "" {
		id, err := strconv.Atoi(vmID)
		if err != nil {
			h.log.Error("invalid vending_machine_id parameter", "error", err)
			http.Error(w, "invalid vending_machine_id parameter", http.StatusBadRequest)
			return
		}
		vendingMachineID = id
	}

	sales, err := h.uc.GetSales(r.Context(), vendingMachineID, fromDate, toDate)
	if err != nil {
		h.log.Error("failed to get sales", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sales); err != nil {
		h.log.Error("failed to encode response", "error", err)
	}
}

func (h *VendingMachineHandler) GetActions(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from_date")
	toDate := r.URL.Query().Get("to_date")
	vendingMachineID := 0

	if vmID := r.URL.Query().Get("vending_machine_id"); vmID != "" {
		id, err := strconv.Atoi(vmID)
		if err != nil {
			h.log.Error("invalid vending_machine_id parameter", "error", err)
			http.Error(w, "invalid vending_machine_id parameter", http.StatusBadRequest)
			return
		}
		vendingMachineID = id
	}

	actions, err := h.uc.GetActions(r.Context(), vendingMachineID, fromDate, toDate)
	if err != nil {
		h.log.Error("failed to get actions", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(actions); err != nil {
		h.log.Error("failed to encode response", "error", err)
	}
}

func (h *VendingMachineHandler) GetVMStates(w http.ResponseWriter, r *http.Request) {
	states, err := h.uc.GetVMStates(r.Context())
	if err != nil {
		h.log.Error("failed to get vm states", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(states); err != nil {
		h.log.Error("failed to encode response", "error", err)
	}
}

func (h *VendingMachineHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from_date")
	toDate := r.URL.Query().Get("to_date")
	vendingMachineID := 0

	if vmID := r.URL.Query().Get("vending_machine_id"); vmID != "" {
		id, err := strconv.Atoi(vmID)
		if err != nil {
			h.log.Error("invalid vending_machine_id parameter", "error", err)
			http.Error(w, "invalid vending_machine_id parameter", http.StatusBadRequest)
			return
		}
		vendingMachineID = id
	}

	events, err := h.uc.GetEvents(r.Context(), vendingMachineID, fromDate, toDate)
	if err != nil {
		h.log.Error("failed to get events", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		h.log.Error("failed to encode response", "error", err)
	}
}

func (h *VendingMachineHandler) SendCommand(w http.ResponseWriter, r *http.Request) {
	vendingMachineID := r.URL.Query().Get("vending_machine_id")
	commandCode := r.URL.Query().Get("command_code")

	if vendingMachineID == "" || commandCode == "" {
		h.log.Error("missing required parameters")
		http.Error(w, "missing vending_machine_id or command_code", http.StatusBadRequest)
		return
	}

	vmID, err := strconv.Atoi(vendingMachineID)
	if err != nil {
		h.log.Error("invalid vending_machine_id", "error", err)
		http.Error(w, "invalid vending_machine_id", http.StatusBadRequest)
		return
	}

	cmdCode, err := strconv.Atoi(commandCode)
	if err != nil {
		h.log.Error("invalid command_code", "error", err)
		http.Error(w, "invalid command_code", http.StatusBadRequest)
		return
	}

	err = h.uc.SendCommand(r.Context(), vmID, cmdCode)
	if err != nil {
		h.log.Error("failed to send command", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "success"}); err != nil {
		h.log.Error("failed to encode response", "error", err)
	}
}

func (h *VendingMachineHandler) GetVendingMachineRemains(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.log.Error("missing id parameter")
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	vendingMachineID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Error("invalid id parameter", "error", err)
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	remains, err := h.uc.GetVendingMachineRemains(r.Context(), vendingMachineID)
	if err != nil {
		h.log.Error("failed to get vending machine remains", "id", vendingMachineID, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(remains); err != nil {
		h.log.Error("failed to encode response", "error", err)
	}
}
