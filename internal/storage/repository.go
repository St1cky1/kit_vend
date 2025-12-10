package storage

import (
	"context"

	"github.com/St1cky1/kit_vend/internal/entity"
)

type VendingMachineRepository interface {
	GetByID(ctx context.Context, id int) (*entity.VendingMachine, error)
	GetAll(ctx context.Context) ([]entity.VendingMachine, error)
	Create(ctx context.Context, vm *entity.VendingMachine) error
	Update(ctx context.Context, vm *entity.VendingMachine) error
	Delete(ctx context.Context, id int) error
}

type SaleRepository interface {
	GetByFilter(ctx context.Context, vendingMachineId int, fromDate, toDate string) ([]entity.Sale, error)
	Create(ctx context.Context, sale *entity.Sale) error
}

type ActionRepository interface {
	GetByFilter(ctx context.Context, vendingMachineId int, fromDate, toDate string) ([]entity.Action, error)
	Create(ctx context.Context, action *entity.Action) error
}

type EventRepository interface {
	GetByFilter(ctx context.Context, vendingMachineId int, fromDate, toDate string) ([]entity.Event, error)
	Create(ctx context.Context, event *entity.Event) error
}

type VMStateRepository interface {
	GetByVendingMachineID(ctx context.Context, id int) (*entity.VMState, error)
	GetAll(ctx context.Context) ([]entity.VMState, error)
}

type VendingMachineRemainsRepository interface {
	GetByVendingMachineID(ctx context.Context, id int) ([]entity.VendingMachineRemains, error)
}
