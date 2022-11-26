package util

import (
	"strings"
)

//used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

//used when data doesn't want to be null on json
type EmptyObj struct{}

//to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, data interface{}) Response {
	splittedError := strings.Split(message, "\n")
	res := Response{
		Status:  false,
		Message: splittedError,
		Data:    nil,
	}

	return res
}
