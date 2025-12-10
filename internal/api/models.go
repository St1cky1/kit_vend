package api

type Auth struct {
	CompanyId int    `json:"CompanyId"`
	RequestId int64  `json:"RequestId"`
	UserLogin string `json:"UserLogin"`
	Sign      string `json:"Sign"`
}

type Company struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type GetCompaniesResponse struct {
	ResultCode int       `json:"ResultCode"`
	Companies  []Company `json:"Companies"`
}

type Modem struct {
	Id              int    `json:"Id"`
	Name            string `json:"Name"`
	VendingMachineId int   `json:"VendingMachineId"`
}

type GetModemsResponse struct {
	ResultCode int     `json:"ResultCode"`
	Modems     []Modem `json:"Modems"`
}

type VendingMachine struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	CompanyId int   `json:"CompanyId"`
}

type GetVendingMachinesResponse struct {
	ResultCode       int              `json:"ResultCode"`
	VendingMachines  []VendingMachine `json:"VendingMachines"`
}

type Filter struct {
	FromDate         string `json:"FromDate"`
	ToDate           string `json:"ToDate"`
	CompanyId        int    `json:"CompanyId,omitempty"`
	VendingMachineId int    `json:"VendingMachineId,omitempty"`
}

type Command struct {
	CommandCode      int `json:"CommandCode"`
	VendingMachineId int `json:"VendingMachineId"`
}

type Sale struct {
	Id               int    `json:"Id"`
	VendingMachineId int    `json:"VendingMachineId"`
	GoodsId          int    `json:"GoodsId"`
	GoodsName        string `json:"GoodsName"`
	Count            int    `json:"Count"`
	Sum              float64 `json:"Sum"`
	DateTime         string `json:"DateTime"`
	PaymentMethod    int    `json:"PaymentMethod"`
}

type GetSalesResponse struct {
	ResultCode int    `json:"ResultCode"`
	Sales      []Sale `json:"Sales"`
}

type Action struct {
	Id               int    `json:"Id"`
	VendingMachineId int    `json:"VendingMachineId"`
	ActionType       int    `json:"ActionType"`
	DateTime         string `json:"DateTime"`
}

type GetActionsResponse struct {
	ResultCode int      `json:"ResultCode"`
	Actions    []Action `json:"Actions"`
}

type Event struct {
	Id               int    `json:"Id"`
	VendingMachineId int    `json:"VendingMachineId"`
	EventCode        int    `json:"EventCode"`
	DateTime         string `json:"DateTime"`
	Description      string `json:"Description"`
}

type GetEventsResponse struct {
	ResultCode int     `json:"ResultCode"`
	Events     []Event `json:"Events"`
}

type VMState struct {
	VendingMachineId int    `json:"VendingMachineId"`
	MachineName      string `json:"MachineName"`
	IsOnline         bool   `json:"IsOnline"`
	LastActivityTime string `json:"LastActivityTime"`
	PowerSupply      int    `json:"PowerSupply"`
	BillAcceptor     int    `json:"BillAcceptor"`
	CoinAcceptor     int    `json:"CoinAcceptor"`
	NonCashPayment   int    `json:"NonCashPayment"`
	CashRegister     int    `json:"CashRegister"`
	QRDisplay        int    `json:"QRDisplay"`
}

type GetVMStatesResponse struct {
	ResultCode int      `json:"ResultCode"`
	VMStates   []VMState `json:"VMStates"`
}

type VendingMachineRemains struct {
	VendingMachineId int    `json:"VendingMachineId"`
	GoodsId          int    `json:"GoodsId"`
	GoodsName        string `json:"GoodsName"`
	Count            int    `json:"Count"`
}

type GetVendingMachineRemainsResponse struct {
	ResultCode int                      `json:"ResultCode"`
	Remains    []VendingMachineRemains `json:"Remains"`
}

type SendCommandResponse struct {
	ResultCode int `json:"ResultCode"`
	CommandId  int `json:"CommandId"`
}

type User struct {
	Id    int    `json:"Id"`
	Login string `json:"Login"`
	Name  string `json:"Name"`
}

type GetUsersResponse struct {
	ResultCode int    `json:"ResultCode"`
	Users      []User `json:"Users"`
}

type Good struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	Code       string `json:"Code"`
	VendorCode string `json:"VendorCode"`
	ShortName  string `json:"ShortName"`
}

type GetGoodsResponse struct {
	ResultCode int    `json:"ResultCode"`
	Goods      []Good `json:"Goods"`
}
