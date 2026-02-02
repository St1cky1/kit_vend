package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/St1cky1/kit_vend/internal/api"
	"github.com/St1cky1/kit_vend/internal/api/kit_vending"
	"github.com/St1cky1/kit_vend/internal/entity"
	"github.com/St1cky1/kit_vend/internal/storage"
)

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrMachineOffline     = errors.New("machine is offline")
	ErrSessionAlreadyExists = errors.New("active load session already exists for this machine")
	ErrSessionNotFound     = errors.New("load session not found")
	ErrSessionNotActive    = errors.New("load session is not active")
)

type CourierUseCase struct {
	userRepo    storage.UserRepository
	cellRepo    storage.CellRepository
	sessionRepo storage.LoadSessionRepository
	vmStateRepo storage.VMStateRepository
	kitClient   kit_vending.Provider
}

func NewCourierUseCase(
	userRepo storage.UserRepository,
	cellRepo storage.CellRepository,
	sessionRepo storage.LoadSessionRepository,
	vmStateRepo storage.VMStateRepository,
	kitClient kit_vending.Provider,
) *CourierUseCase {
	return &CourierUseCase{
		userRepo:    userRepo,
		cellRepo:    cellRepo,
		sessionRepo: sessionRepo,
		vmStateRepo: vmStateRepo,
		kitClient:   kitClient,
	}
}

func (uc *CourierUseCase) checkRole(ctx context.Context, userID int, requiredRole string) error {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return ErrUnauthorized
	}
	if user.Role != requiredRole {
		return ErrForbidden
	}
	return nil
}

func (uc *CourierUseCase) GetCells(ctx context.Context, courierID int, vmID int) ([]entity.Cell, error) {
	if err := uc.checkRole(ctx, courierID, entity.RoleCourier); err != nil {
		return nil, err
	}
	return uc.cellRepo.GetByVendingMachineID(ctx, vmID)
}

func (uc *CourierUseCase) StartLoadSession(ctx context.Context, courierID int, vmID int) (*entity.LoadSession, error) {
	if err := uc.checkRole(ctx, courierID, entity.RoleCourier); err != nil {
		return nil, err
	}

	// Проверка статуса автомата
	state, err := uc.vmStateRepo.GetByVendingMachineID(ctx, vmID)
	if err != nil {
		return nil, fmt.Errorf("failed to get machine state: %w", err)
	}
	if !state.IsOnline {
		return nil, ErrMachineOffline
	}

	// Проверка наличия активной сессии
	existing, err := uc.sessionRepo.GetActiveByVendingMachineID(ctx, vmID)
	if err == nil && existing != nil {
		return nil, ErrSessionAlreadyExists
	}

	session := &entity.LoadSession{
		VendingMachineId: vmID,
		CourierId:        courierID,
		Status:           entity.LoadSessionStatusActive,
		StartedAt:        time.Now(),
	}

	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (uc *CourierUseCase) LoadCell(ctx context.Context, courierID int, sessionID int, cellID string, goodsID int) error {
	if err := uc.checkRole(ctx, courierID, entity.RoleCourier); err != nil {
		return err
	}

	session, err := uc.sessionRepo.GetByID(ctx, sessionID)
	if err != nil {
		return ErrSessionNotFound
	}
	if session.Status != entity.LoadSessionStatusActive {
		return ErrSessionNotActive
	}
	if session.CourierId != courierID {
		return ErrForbidden
	}

	cell := &entity.Cell{
		Id:               cellID,
		VendingMachineId: session.VendingMachineId,
		Status:           entity.CellStatusOccupied,
		GoodsId:          &goodsID,
	}

	return uc.cellRepo.Update(ctx, cell)
}

func (uc *CourierUseCase) CompleteLoadSession(ctx context.Context, courierID int, sessionID int) error {
	if err := uc.checkRole(ctx, courierID, entity.RoleCourier); err != nil {
		return err
	}

	session, err := uc.sessionRepo.GetByID(ctx, sessionID)
	if err != nil {
		return ErrSessionNotFound
	}
	if session.Status != entity.LoadSessionStatusActive {
		return ErrSessionNotActive
	}
	if session.CourierId != courierID {
		return ErrForbidden
	}

	now := time.Now()
	session.Status = entity.LoadSessionStatusCompleted
	session.CompletedAt = &now

	if err := uc.sessionRepo.Update(ctx, session); err != nil {
		return err
	}

	// Отправка команды в Kit Vending API
	command := api.Command{
		CommandCode:      26, // Сделать загрузку
		VendingMachineId: session.VendingMachineId,
	}

	var result api.SendCommandResponse
	err = uc.kitClient.Call("SendCommand", map[string]interface{}{"Command": command}, &result)
	if err != nil {
		return fmt.Errorf("failed to send command to Kit Vending: %w", err)
	}

	return api.CheckResultCode(result.ResultCode)
}
