package response

import "net/http"

// Response 统一返回结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 返回信息
	Data    interface{} `json:"data"`    // 返回数据
}

// Success 成功返回
func Success(data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

// Error 错误返回
func Error(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
