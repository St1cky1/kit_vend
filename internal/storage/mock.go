package storage

import (
	"context"

	"github.com/St1cky1/kit_vend/internal/entity"
)

type MockVendingMachineRepository struct {
	machines []entity.VendingMachine
}

func NewMockVendingMachineRepository() *MockVendingMachineRepository {
	return &MockVendingMachineRepository{
		machines: []entity.VendingMachine{
			{Id: 1, Name: "Machine 1", CompanyId: 380649},
			{Id: 2, Name: "Machine 2", CompanyId: 380649},
		},
	}
}

func (r *MockVendingMachineRepository) GetByID(ctx context.Context, id int) (*entity.VendingMachine, error) {
	for _, m := range r.machines {
		if m.Id == id {
			return &m, nil
		}
	}
	return nil, nil
}

func (r *MockVendingMachineRepository) GetAll(ctx context.Context) ([]entity.VendingMachine, error) {
	return r.machines, nil
}

func (r *MockVendingMachineRepository) Create(ctx context.Context, vm *entity.VendingMachine) error {
	r.machines = append(r.machines, *vm)
	return nil
}

func (r *MockVendingMachineRepository) Update(ctx context.Context, vm *entity.VendingMachine) error {
	for i, m := range r.machines {
		if m.Id == vm.Id {
			r.machines[i] = *vm
			return nil
		}
	}
	return nil
}

func (r *MockVendingMachineRepository) Delete(ctx context.Context, id int) error {
	for i, m := range r.machines {
		if m.Id == id {
			r.machines = append(r.machines[:i], r.machines[i+1:]...)
			return nil
		}
	}
	return nil
}

type MockSaleRepository struct {
	sales []entity.Sale
}

func NewMockSaleRepository() *MockSaleRepository {
	return &MockSaleRepository{
		sales: []entity.Sale{},
	}
}

func (r *MockSaleRepository) GetByFilter(ctx context.Context, vendingMachineId int, fromDate, toDate string) ([]entity.Sale, error) {
	return r.sales, nil
}

func (r *MockSaleRepository) Create(ctx context.Context, sale *entity.Sale) error {
	r.sales = append(r.sales, *sale)
	return nil
}

type MockActionRepository struct {
	actions []entity.Action
}

func NewMockActionRepository() *MockActionRepository {
	return &MockActionRepository{
		actions: []entity.Action{},
	}
}

func (r *MockActionRepository) GetByFilter(ctx context.Context, vendingMachineId int, fromDate, toDate string) ([]entity.Action, error) {
	return r.actions, nil
}

func (r *MockActionRepository) Create(ctx context.Context, action *entity.Action) error {
	r.actions = append(r.actions, *action)
	return nil
}

type MockEventRepository struct {
	events []entity.Event
}

func NewMockEventRepository() *MockEventRepository {
	return &MockEventRepository{
		events: []entity.Event{},
	}
}

func (r *MockEventRepository) GetByFilter(ctx context.Context, vendingMachineId int, fromDate, toDate string) ([]entity.Event, error) {
	return r.events, nil
}

func (r *MockEventRepository) Create(ctx context.Context, event *entity.Event) error {
	r.events = append(r.events, *event)
	return nil
}

type MockVMStateRepository struct {
	states []entity.VMState
}

func NewMockVMStateRepository() *MockVMStateRepository {
	return &MockVMStateRepository{
		states: []entity.VMState{},
	}
}

func (r *MockVMStateRepository) GetByVendingMachineID(ctx context.Context, id int) (*entity.VMState, error) {
	for _, s := range r.states {
		if s.VendingMachineId == id {
			return &s, nil
		}
	}
	return nil, nil
}

func (r *MockVMStateRepository) GetAll(ctx context.Context) ([]entity.VMState, error) {
	return r.states, nil
}

type MockVendingMachineRemainsRepository struct {
	remains []entity.VendingMachineRemains
}

func NewMockVendingMachineRemainsRepository() *MockVendingMachineRemainsRepository {
	return &MockVendingMachineRemainsRepository{
		remains: []entity.VendingMachineRemains{},
	}
}

func (r *MockVendingMachineRemainsRepository) GetByVendingMachineID(ctx context.Context, id int) ([]entity.VendingMachineRemains, error) {
	return r.remains, nil
}
