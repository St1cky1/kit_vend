package kit_vending

import (
	"encoding/json"
	"fmt"
)

type MockClient struct {
	Debug bool
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (m *MockClient) SetDebug(debug bool) {
	m.Debug = debug
}

func (m *MockClient) Call(method string, extra map[string]interface{}, v interface{}) error {
	if m.Debug {
		fmt.Printf("[MOCK DEBUG] Kit Vending API Call: %s\n", method)
	}

	// Формируем базовый успешный ответ
	responseJSON := `{"ResultCode": 0}`

	// Для специфичных методов можно добавить больше данных в JSON
	switch method {
	case "SendCommand":
		responseJSON = `{"ResultCode": 0, "CommandId": 12345}`
	case "GetSales":
		responseJSON = `{"ResultCode": 0, "Sales": []}`
	case "GetVMStates":
		responseJSON = `{"ResultCode": 0, "VMStates": []}`
	}

	if err := json.Unmarshal([]byte(responseJSON), v); err != nil {
		return fmt.Errorf("mock unmarshal error: %w", err)
	}

	return nil
}
