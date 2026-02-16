package grpc

import (
	"errors"

	"github.com/St1cky1/kit_vend/internal/api"
	"github.com/St1cky1/kit_vend/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapError(err error) error {
	if err == nil {
		return nil
	}

	var apiErr *api.APIError
	if errors.As(err, &apiErr) {
		// Используем InvalidArgument или Internal в зависимости от кода, 
		// но главное - передаем сообщение из apiErr
		return status.Error(codes.InvalidArgument, apiErr.Message)
	}

	if errors.Is(err, usecase.ErrUnauthorized) {
		return status.Error(codes.Unauthenticated, err.Error())
	}
	if errors.Is(err, usecase.ErrForbidden) {
		return status.Error(codes.PermissionDenied, err.Error())
	}
	if errors.Is(err, usecase.ErrSessionNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, usecase.ErrMachineOffline) {
		return status.Error(codes.FailedPrecondition, err.Error())
	}
	if errors.Is(err, usecase.ErrSessionAlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	if errors.Is(err, usecase.ErrSessionNotActive) {
		return status.Error(codes.FailedPrecondition, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
