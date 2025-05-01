package presenters

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Details any    `json:"details,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func JSONError(w http.ResponseWriter, statusCode int, message string, code int, details any) {
	resp := ErrorResponse{
		Message: message,
		Code:    code,
		Details: details,
	}
	JSON(w, statusCode, resp)
}
