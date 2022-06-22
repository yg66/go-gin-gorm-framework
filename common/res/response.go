package res

import (
	"github.com/yg66/go-gin-gorm-framework/common/errors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Response struct {
	Ts        int64       `json:"ts"`
	Code      int         `json:"code"`
	ErrorCode int         `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func ErrCodeText(errCode int) (string, bool) {
	message, ok := errors.ErrCodeText[errCode]
	return message, ok
}

func UnknownErr(err error) *Response {
	message, ok := ErrCodeText(errors.UnknownError)
	if !ok {
		zap.S().Errorf("Error {%v}", err)
		message = "System Error"
	}
	return &Response{
		Ts:        time.Now().UTC().Unix(),
		Code:      http.StatusOK,
		ErrorCode: errors.UnknownError,
		Message:   message,
		Data:      nil,
	}
}

func Failed(_errCode int) *Response {
	message, ok := ErrCodeText(_errCode)
	if !ok {
		message = "System Error"
	}
	response := Response{
		Ts:        time.Now().UTC().Unix(),
		Code:      http.StatusOK,
		ErrorCode: _errCode,
		Message:   message,
		Data:      nil,
	}
	return &response
}

func Success(_data interface{}) *Response {
	return &Response{
		Ts:        time.Now().UTC().Unix(),
		Code:      http.StatusOK,
		ErrorCode: 0,
		Message:   "",
		Data:      _data,
	}
}
