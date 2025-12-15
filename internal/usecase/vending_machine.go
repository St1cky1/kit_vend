package usecase

import (
	"context"

	"github.com/St1cky1/kit_vend/internal/api"
	"github.com/St1cky1/kit_vend/internal/entity"
	"github.com/St1cky1/kit_vend/internal/storage"
)

type VendingMachineUseCase struct {
	kitClient   *api.Client
	vmRepo      storage.VendingMachineRepository
	saleRepo    storage.SaleRepository
	actionRepo  storage.ActionRepository
	eventRepo   storage.EventRepository
	vmStateRepo storage.VMStateRepository
	remainsRepo storage.VendingMachineRemainsRepository
}

func NewVendingMachineUseCase(
	kitClient *api.Client,
	vmRepo storage.VendingMachineRepository,
	saleRepo storage.SaleRepository,
	actionRepo storage.ActionRepository,
	eventRepo storage.EventRepository,
	vmStateRepo storage.VMStateRepository,
	remainsRepo storage.VendingMachineRemainsRepository,
) *VendingMachineUseCase {
	return &VendingMachineUseCase{
		kitClient:   kitClient,
		vmRepo:      vmRepo,
		saleRepo:    saleRepo,
		actionRepo:  actionRepo,
		eventRepo:   eventRepo,
		vmStateRepo: vmStateRepo,
		remainsRepo: remainsRepo,
	}
}

func (uc *VendingMachineUseCase) GetVendingMachineByID(ctx context.Context, id int) (*entity.VendingMachine, error) {
	return uc.vmRepo.GetByID(ctx, id)
}

func (uc *VendingMachineUseCase) GetSales(ctx context.Context, vendingMachineId int, upDate, toDate string) ([]entity.Sale, error) {
	filter := api.Filter{
		UpDate: upDate,
		ToDate: toDate,
	}
	if vendingMachineId > 0 {
		filter.VendingMachineId = vendingMachineId
	}

	var result api.GetSalesResponse
	err := uc.kitClient.Call("GetSales", map[string]interface{}{"Filter": filter}, &result)
	if err != nil {
		return nil, err
	}

	sales := make([]entity.Sale, len(result.Sales))
	for i, s := range result.Sales {
		sales[i] = entity.Sale{
			Id:               s.Id,
			VendingMachineId: s.VendingMachineId,
			GoodsId:          s.GoodsId,
			GoodsName:        s.GoodsName,
			Count:            s.Count,
			Sum:              s.Sum,
			DateTime:         s.DateTime,
			PaymentMethod:    s.PaymentMethod,
		}
	}

	return sales, nil
}

func (uc *VendingMachineUseCase) GetActions(ctx context.Context, vendingMachineId int, upDate, toDate string) ([]entity.Action, error) {
	filter := api.Filter{
		UpDate: upDate,
		ToDate: toDate,
	}
	if vendingMachineId > 0 {
		filter.VendingMachineId = vendingMachineId
	}

	var result api.GetActionsResponse
	err := uc.kitClient.Call("GetActions", map[string]interface{}{"Filter": filter}, &result)
	if err != nil {
		return nil, err
	}

	actions := make([]entity.Action, len(result.Actions))
	for i, a := range result.Actions {
		actions[i] = entity.Action{
			Id:               a.Id,
			VendingMachineId: a.VendingMachineId,
			ActionType:       a.ActionType,
			DateTime:         a.DateTime,
		}
	}

	return actions, nil
}

func (uc *VendingMachineUseCase) GetVMStates(ctx context.Context) ([]entity.VMState, error) {
	var result api.GetVMStatesResponse
	err := uc.kitClient.Call("GetVMStates", nil, &result)
	if err != nil {
		return nil, err
	}

	states := make([]entity.VMState, len(result.VMStates))
	for i, s := range result.VMStates {
		states[i] = entity.VMState{
			VendingMachineId: s.VendingMachineId,
			MachineName:      s.MachineName,
			IsOnline:         s.IsOnline,
			LastActivityTime: s.LastActivityTime,
			PowerSupply:      s.PowerSupply,
			BillAcceptor:     s.BillAcceptor,
			CoinAcceptor:     s.CoinAcceptor,
			NonCashPayment:   s.NonCashPayment,
			CashRegister:     s.CashRegister,
			QRDisplay:        s.QRDisplay,
		}
	}

	return states, nil
}

func (uc *VendingMachineUseCase) GetEvents(ctx context.Context, vendingMachineId int, upDate, toDate string) ([]entity.Event, error) {
	filter := api.Filter{
		UpDate: upDate,
		ToDate: toDate,
	}
	if vendingMachineId > 0 {
		filter.VendingMachineId = vendingMachineId
	}

	var result api.GetEventsResponse
	err := uc.kitClient.Call("GetEvents", map[string]interface{}{"Filter": filter}, &result)
	if err != nil {
		return nil, err
	}

	events := make([]entity.Event, len(result.Events))
	for i, e := range result.Events {
		events[i] = entity.Event{
			Id:               e.Id,
			VendingMachineId: e.VendingMachineId,
			EventCode:        e.EventCode,
			DateTime:         e.DateTime,
			Description:      e.Description,
		}
	}

	return events, nil
}

func (uc *VendingMachineUseCase) SendCommand(ctx context.Context, vendingMachineId int, commandCode int) error {
	command := api.Command{
		CommandCode:      commandCode,
		VendingMachineId: vendingMachineId,
	}

	var result api.SendCommandResponse
	err := uc.kitClient.Call("SendCommand", map[string]interface{}{"Command": command}, &result)
	if err != nil {
		return err
	}

	return nil
}

func (uc *VendingMachineUseCase) GetVendingMachineRemains(ctx context.Context, id int) ([]entity.VendingMachineRemains, error) {
	var result api.GetVendingMachineRemainsResponse
	err := uc.kitClient.Call("GetVendingMachineRemains", map[string]interface{}{"Id": id}, &result)
	if err != nil {
		return nil, err
	}

	remains := make([]entity.VendingMachineRemains, len(result.Remains))
	for i, r := range result.Remains {
		remains[i] = entity.VendingMachineRemains{
			VendingMachineId: r.VendingMachineId,
			GoodsId:          r.GoodsId,
			GoodsName:        r.GoodsName,
			Count:            r.Count,
		}
	}

	return remains, nil
}
