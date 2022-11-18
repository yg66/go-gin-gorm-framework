package errors

import (
	"encoding/json"
	"fmt"
)

type Err struct {
	Code    int
	Message string
}

const (
	ServerError                   = 1000
	NotFound                      = 1001
	UnknownError                  = 1002
	ParameterError                = 1003
	NetworkAnomaly                = 1004
	OperationTooFrequently        = 1005
	UriNotFoundOrMethodNotSupport = 1404
	PortNotOpen                   = 1500
)

var ErrCodeText = map[int]string{
	ServerError:                   "Server Error",
	NotFound:                      "Not Found",
	UnknownError:                  "Unknown Error",
	ParameterError:                "Parameter Error",
	NetworkAnomaly:                "Network Anomaly",
	OperationTooFrequently:        "Operation too frequently",
	UriNotFoundOrMethodNotSupport: "Uri not found or method can not support",
	PortNotOpen:                   "The port :%d not open",
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

// New
// a ...interface{}: Indicates a placeholder
func New(code int, a ...interface{}) *Err {
	message := ErrCodeText[code]
	if len(a) > 0 {
		message = fmt.Errorf(message, a).Error()
	}
	return &Err{
		Code:    code,
		Message: message,
	}
}
