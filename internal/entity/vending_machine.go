package entity

type VendingMachine struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CompanyId int    `json:"companyId"`
}

type Sale struct {
	Id               int     `json:"id"`
	VendingMachineId int     `json:"vendingMachineId"`
	GoodsId          int     `json:"goodsId"`
	GoodsName        string  `json:"goodsName"`
	Count            int     `json:"count"`
	Sum              float64 `json:"sum"`
	DateTime         string  `json:"dateTime"`
	PaymentMethod    int     `json:"paymentMethod"`
}

type Action struct {
	Id               int    `json:"id"`
	VendingMachineId int    `json:"vendingMachineId"`
	ActionType       int    `json:"actionType"`
	DateTime         string `json:"dateTime"`
}

type Event struct {
	Id               int    `json:"id"`
	VendingMachineId int    `json:"vendingMachineId"`
	EventCode        int    `json:"eventCode"`
	DateTime         string `json:"dateTime"`
	Description      string `json:"description"`
}

type VMState struct {
	VendingMachineId int    `json:"vendingMachineId"`
	MachineName      string `json:"machineName"`
	IsOnline         bool   `json:"isOnline"`
	LastActivityTime string `json:"lastActivityTime"`
	PowerSupply      int    `json:"powerSupply"`
	BillAcceptor     int    `json:"billAcceptor"`
	CoinAcceptor     int    `json:"coinAcceptor"`
	NonCashPayment   int    `json:"nonCashPayment"`
	CashRegister     int    `json:"cashRegister"`
	QRDisplay        int    `json:"qrDisplay"`
}

type VendingMachineRemains struct {
	VendingMachineId int    `json:"vendingMachineId"`
	GoodsId          int    `json:"goodsId"`
	GoodsName        string `json:"goodsName"`
	Count            int    `json:"count"`
}
