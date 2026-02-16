package kit_vending

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	//"github.com/St1cky1/kit_vend/internal/api"

	"github.com/St1cky1/kit_vend/internal/api"
)

const baseUrl = "https://api2.kit-invest.ru/APIService.svc/"

type Provider interface {
	Call(method string, extra map[string]interface{}, v interface{}) error
	SetDebug(debug bool)
}

type Client struct {
	CompanyId int
	Login     string
	Password  string
	Http      *http.Client
	Debug     bool
}

func NewClient(companyId int, login, password string) *Client {
	return &Client{
		CompanyId: companyId,
		Login:     login,
		Password:  password,
		Http:      &http.Client{Timeout: 10 * time.Second},
		Debug:     false,
	}
}

func (c *Client) SetDebug(debug bool) {
	c.Debug = debug
}

// генерируем уникальные RequestID в формате YYMMDDHHMMSS
func (c *Client) generateRequestid() int64 {
	s := time.Now().Format("060102150405")
	val, _ := strconv.ParseInt(s, 10, 64)
	return val
}

// описываем функцию, которая возвращает значение функции MD5
func (c *Client) makeSign(requestId int64) string {
	raw := fmt.Sprintf("%d%s%d", c.CompanyId, c.Password, requestId)
	hash := md5.Sum([]byte(raw))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

func (c *Client) autPayload() api.Auth {
	reqId := c.generateRequestid()
	return api.Auth{
		CompanyId: c.CompanyId,
		RequestId: reqId,
		UserLogin: c.Login,
		Sign:      c.makeSign(reqId),
	}
}

// генериуем запрос к API и получаем ответ
func (c *Client) Call(method string, extra map[string]interface{}, v interface{}) error {
	url := baseUrl + method

	payload := map[string]interface{}{
		"Auth": c.autPayload(),
	}
	for k, v := range extra {
		payload[k] = v
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	if c.Debug {
		fmt.Printf("[DEBUG] Kit Vending API Request:\nURL: %s\nBody: %s\n", url, string(body))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.Http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("http error: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}

	if c.Debug {
		fmt.Printf("[DEBUG] Kit Vending API Response:\n%s\n", string(respBody))
	}

	// Сначала проверяем ResultCode
	var baseResp struct {
		ResultCode int `json:"ResultCode"`
	}
	if err := json.Unmarshal(respBody, &baseResp); err == nil {
		if err := api.CheckResponse(baseResp.ResultCode); err != nil {
			return err
		}
	}

	if err := json.Unmarshal(respBody, &v); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}

	return nil
}
