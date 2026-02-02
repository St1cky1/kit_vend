package grpc

import (
	"context"

	"github.com/St1cky1/kit_vend/internal/usecase"
	pbv1 "github.com/St1cky1/kit_vend/pkg/pb1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CourierServiceServer struct {
	pbv1.UnimplementedCourierServiceServer
	uc *usecase.CourierUseCase
}

func NewCourierServiceServer(uc *usecase.CourierUseCase) *CourierServiceServer {
	return &CourierServiceServer{
		uc: uc,
	}
}

func (s *CourierServiceServer) GetCells(ctx context.Context, req *pbv1.GetCellsRequest) (*pbv1.GetCellsResponse, error) {
	// В реальном приложении courier_id должен браться из контекста (после аутентификации)
	// Для MVP передадим хардкодом или добавим в запрос, если нужно. 
	// Но по ТЗ курьер авторизуется, значит ID должен быть.
	// В данном MVP предположим, что мы можем получить его из метаданных или просто использовать тестовый ID.
	courierID := 1 // Mock courier ID

	cells, err := s.uc.GetCells(ctx, courierID, int(req.VendingMachineId))
	if err != nil {
		return nil, err
	}

	protoCells := make([]*pbv1.Cell, len(cells))
	for i, c := range cells {
		protoCells[i] = &pbv1.Cell{
			Id:               c.Id,
			VendingMachineId: int32(c.VendingMachineId),
			Status:           string(c.Status),
		}
		if c.GoodsId != nil {
			goodsID := int32(*c.GoodsId)
			protoCells[i].GoodsId = &goodsID
		}
	}

	return &pbv1.GetCellsResponse{
		Cells: protoCells,
	}, nil
}

func (s *CourierServiceServer) StartLoadSession(ctx context.Context, req *pbv1.StartLoadSessionRequest) (*pbv1.LoadSession, error) {
	session, err := s.uc.StartLoadSession(ctx, int(req.CourierId), int(req.VendingMachineId))
	if err != nil {
		return nil, err
	}

	res := &pbv1.LoadSession{
		Id:               int32(session.Id),
		VendingMachineId: int32(session.VendingMachineId),
		CourierId:        int32(session.CourierId),
		Status:           string(session.Status),
		StartedAt:        timestamppb.New(session.StartedAt),
	}
	if session.CompletedAt != nil {
		res.CompletedAt = timestamppb.New(*session.CompletedAt)
	}

	return res, nil
}

func (s *CourierServiceServer) LoadCell(ctx context.Context, req *pbv1.LoadCellRequest) (*emptypb.Empty, error) {
	// Mock courier ID
	courierID := 1 

	err := s.uc.LoadCell(ctx, courierID, int(req.SessionId), req.CellId, int(req.GoodsId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *CourierServiceServer) CompleteLoadSession(ctx context.Context, req *pbv1.CompleteLoadSessionRequest) (*emptypb.Empty, error) {
	// Mock courier ID
	courierID := 1

	err := s.uc.CompleteLoadSession(ctx, courierID, int(req.SessionId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
