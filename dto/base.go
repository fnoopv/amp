package dto

// CommonResponse 基础响应结构
type CommonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

var (
	// SuccessMessage 成功且无需返回数据
	SuccessResponse = &CommonResponse{
		Message: SuccessMessage,
		Data:    nil,
	}
	// SuccessMessage 默认成功消息
	SuccessMessage = "success"
)
