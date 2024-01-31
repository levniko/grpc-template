package models

type Response struct {
	Message   string `json:"message"`
	Error     bool   `json:"error"`
	ErrorCode int64  `json:"error_code"`
	Data      any    `json:"data,omitempty"`
}
