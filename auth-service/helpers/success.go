package helpers

import "net/http"

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"statusCode,omitempty"`
	Message string      `json:"message,omitempty"`
}

func Success(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
	}
}
