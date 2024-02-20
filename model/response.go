package model

import "strings"

type Response struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func BuildResponse(message string, error string, data interface{}) Response {
	if len(error) == 0 {
		return Response{
			Message: message,
			Errors:  nil,
			Data:    data,
		}
	}
	errors := strings.Split(error, "\n")
	res := Response{
		Message: message,
		Errors:  errors,
		Data:    data,
	}
	return res
}
