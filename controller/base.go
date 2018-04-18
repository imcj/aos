package controller

type ResponseObject struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Base 基类
type Base struct {
}
