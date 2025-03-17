package dto

// CommonResponse 基础响应结构
type CommonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

var (
	// ResponseSuccess 成功且无需返回数据
	ResponseSuccess = &CommonResponse{
		Message: ResponseSuccessMessage,
		Data:    nil,
	}
	// ResponseSuccessMessage 默认成功消息
	ResponseSuccessMessage = "success"
)
