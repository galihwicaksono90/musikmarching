package response

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponseError(code int, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

func HandleError(w http.ResponseWriter, code int, message string) {
	response := NewResponseError(code, message)
	json.NewEncoder(w).Encode(response)
}
