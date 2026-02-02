package entity

import "time"

const (
	RoleAdmin   = "admin"
	RoleCourier = "courier"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type CellStatus string

const (
	CellStatusFree     CellStatus = "free"
	CellStatusOccupied CellStatus = "occupied"
	CellStatusBlocked  CellStatus = "blocked"
)

type Cell struct {
	Id               string     `json:"id"`
	VendingMachineId int        `json:"vendingMachineId"`
	Status           CellStatus `json:"status"`
	GoodsId          *int       `json:"goodsId,omitempty"`
}

type LoadSessionStatus string

const (
	LoadSessionStatusActive    LoadSessionStatus = "active"
	LoadSessionStatusCompleted LoadSessionStatus = "completed"
	LoadSessionStatusCancelled LoadSessionStatus = "cancelled"
)

type LoadSession struct {
	Id               int               `json:"id"`
	VendingMachineId int               `json:"vendingMachineId"`
	CourierId        int               `json:"courierId"`
	Status           LoadSessionStatus `json:"status"`
	StartedAt        time.Time         `json:"startedAt"`
	CompletedAt      *time.Time        `json:"completedAt,omitempty"`
}
