package model

// RespondStruct 请求返回的通用结构体
type RespondStruct struct {
	Status  string      `json:"status"`
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *string     `json:"error,omitempty"`
}
