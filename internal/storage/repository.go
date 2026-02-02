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

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*entity.User, error)
}

type CellRepository interface {
	GetByVendingMachineID(ctx context.Context, vmID int) ([]entity.Cell, error)
	Update(ctx context.Context, cell *entity.Cell) error
}

type LoadSessionRepository interface {
	Create(ctx context.Context, session *entity.LoadSession) error
	GetByID(ctx context.Context, id int) (*entity.LoadSession, error)
	GetActiveByVendingMachineID(ctx context.Context, vmID int) (*entity.LoadSession, error)
	Update(ctx context.Context, session *entity.LoadSession) error
}
