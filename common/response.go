package common

import "net/http"

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResponse(message string, data any) BaseResponse {
	return BaseResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(errorStatus int, message string, data any) BaseResponse {
	return BaseResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
}
