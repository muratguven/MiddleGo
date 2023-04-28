package utils

import "github.com/gofiber/fiber/v2"

type BaseResponseModel struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

//BaseResponseModel helpers

// custom response
func NewResponse(status int, data interface{}, message string, success bool) BaseResponseModel {
	return BaseResponseModel{
		Status:  status,
		Data:    data,
		Message: message,
		Success: success,
	}
}

// returns http 200 OK
func StatusOK(data interface{}) BaseResponseModel {
	return BaseResponseModel{
		Status:  fiber.StatusOK,
		Data:    data,
		Message: "OK",
		Success: true,
	}
}

// returns http 400
func StatusFail(message string) BaseResponseModel {
	return BaseResponseModel{
		Status:  fiber.StatusBadRequest,
		Message: message,
		Success: false,
	}
}

// retuns http 401
func StatusUnauthorized(message string) BaseResponseModel {
	return BaseResponseModel{
		Status:  fiber.StatusUnauthorized,
		Message: message,
		Success: false,
	}
}

// returns http 500
func UnhandledError() BaseResponseModel {
	return BaseResponseModel{
		Status:  fiber.StatusInternalServerError,
		Message: "Unhandled error occurred. Please try again later",
		Success: false,
	}
}

// returns http 404
func StatusNotFound(message string) BaseResponseModel {
	return BaseResponseModel{
		Status:  fiber.StatusNotFound,
		Message: message,
		Success: false,
	}
}
