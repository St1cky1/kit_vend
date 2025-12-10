package grpc

import (
	"context"

	apiv1 "github.com/St1cky1/kit_vend/api/v1"
	"github.com/St1cky1/kit_vend/internal/usecase"
)

type VendingMachineServiceServer struct {
	apiv1.UnimplementedVendingMachineServiceServer
	uc *usecase.VendingMachineUseCase
}

func NewVendingMachineServiceServer(uc *usecase.VendingMachineUseCase) *VendingMachineServiceServer {
	return &VendingMachineServiceServer{
		uc: uc,
	}
}

func (s *VendingMachineServiceServer) GetVendingMachineByID(ctx context.Context, req *apiv1.GetVendingMachineByIDRequest) (*apiv1.GetVendingMachineByIDResponse, error) {
	vm, err := s.uc.GetVendingMachineByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &apiv1.GetVendingMachineByIDResponse{
		VendingMachine: &apiv1.VendingMachine{
			Id:        int32(vm.Id),
			Name:      vm.Name,
			CompanyId: int32(vm.CompanyId),
		},
	}, nil
}

func (s *VendingMachineServiceServer) GetSales(ctx context.Context, req *apiv1.GetSalesRequest) (*apiv1.GetSalesResponse, error) {
	sales, err := s.uc.GetSales(ctx, int(req.VendingMachineId), req.FromDate, req.ToDate)
	if err != nil {
		return nil, err
	}

	protoSales := make([]*apiv1.Sale, len(sales))
	for i, s := range sales {
		protoSales[i] = &apiv1.Sale{
			Id:               int32(s.Id),
			VendingMachineId: int32(s.VendingMachineId),
			GoodsId:          int32(s.GoodsId),
			GoodsName:        s.GoodsName,
			Count:            int32(s.Count),
			Sum:              s.Sum,
			DateTime:         s.DateTime,
			PaymentMethod:    int32(s.PaymentMethod),
		}
	}

	return &apiv1.GetSalesResponse{
		Sales: protoSales,
	}, nil
}

func (s *VendingMachineServiceServer) GetActions(ctx context.Context, req *apiv1.GetActionsRequest) (*apiv1.GetActionsResponse, error) {
	actions, err := s.uc.GetActions(ctx, int(req.VendingMachineId), req.FromDate, req.ToDate)
	if err != nil {
		return nil, err
	}

	protoActions := make([]*apiv1.Action, len(actions))
	for i, a := range actions {
		protoActions[i] = &apiv1.Action{
			Id:               int32(a.Id),
			VendingMachineId: int32(a.VendingMachineId),
			ActionType:       int32(a.ActionType),
			DateTime:         a.DateTime,
		}
	}

	return &apiv1.GetActionsResponse{
		Actions: protoActions,
	}, nil
}

func (s *VendingMachineServiceServer) GetVMStates(ctx context.Context, req *apiv1.GetVMStatesRequest) (*apiv1.GetVMStatesResponse, error) {
	states, err := s.uc.GetVMStates(ctx)
	if err != nil {
		return nil, err
	}

	protoStates := make([]*apiv1.VMState, len(states))
	for i, st := range states {
		protoStates[i] = &apiv1.VMState{
			VendingMachineId: int32(st.VendingMachineId),
			MachineName:      st.MachineName,
			IsOnline:         st.IsOnline,
			LastActivityTime: st.LastActivityTime,
			PowerSupply:      int32(st.PowerSupply),
			BillAcceptor:     int32(st.BillAcceptor),
			CoinAcceptor:     int32(st.CoinAcceptor),
			NonCashPayment:   int32(st.NonCashPayment),
			CashRegister:     int32(st.CashRegister),
			QrDisplay:        int32(st.QRDisplay),
		}
	}

	return &apiv1.GetVMStatesResponse{
		VmStates: protoStates,
	}, nil
}

func (s *VendingMachineServiceServer) GetEvents(ctx context.Context, req *apiv1.GetEventsRequest) (*apiv1.GetEventsResponse, error) {
	events, err := s.uc.GetEvents(ctx, int(req.VendingMachineId), req.FromDate, req.ToDate)
	if err != nil {
		return nil, err
	}

	protoEvents := make([]*apiv1.Event, len(events))
	for i, e := range events {
		protoEvents[i] = &apiv1.Event{
			Id:               int32(e.Id),
			VendingMachineId: int32(e.VendingMachineId),
			EventCode:        int32(e.EventCode),
			DateTime:         e.DateTime,
			Description:      e.Description,
		}
	}

	return &apiv1.GetEventsResponse{
		Events: protoEvents,
	}, nil
}

func (s *VendingMachineServiceServer) SendCommand(ctx context.Context, req *apiv1.SendCommandRequest) (*apiv1.SendCommandResponse, error) {
	err := s.uc.SendCommand(ctx, int(req.Command.VendingMachineId), int(req.Command.CommandCode))
	if err != nil {
		return nil, err
	}

	return &apiv1.SendCommandResponse{
		ResultCode: 0,
		CommandId:  1,
	}, nil
}

func (s *VendingMachineServiceServer) GetVendingMachineRemains(ctx context.Context, req *apiv1.GetVendingMachineRemainsRequest) (*apiv1.GetVendingMachineRemainsResponse, error) {
	remains, err := s.uc.GetVendingMachineRemains(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	protoRemains := make([]*apiv1.VendingMachineRemains, len(remains))
	for i, r := range remains {
		protoRemains[i] = &apiv1.VendingMachineRemains{
			VendingMachineId: int32(r.VendingMachineId),
			GoodsId:          int32(r.GoodsId),
			GoodsName:        r.GoodsName,
			Count:            int32(r.Count),
		}
	}

	return &apiv1.GetVendingMachineRemainsResponse{
		Remains: protoRemains,
	}, nil
}
