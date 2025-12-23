package api

import (
	"fmt"

	"github.com/St1cky1/kit_vend/pkg/constants"
)

type APIError struct {
	Code    constants.ErrorCode
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error code %d: %s", e.Code, e.Message)
}

func CheckResultCode(code int) error {
	errorCode := constants.ErrorCode(code)

	if errorCode == constants.ResultCodeSuccess {
		return nil
	}

	message := errorCode.String()
	return &APIError{
		Code:    errorCode,
		Message: message,
	}
}

func CheckResponse(resultCode int) error {
	return CheckResultCode(resultCode)
}
