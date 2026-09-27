package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, status int, success bool, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ApiResponse{
		Success: success,
		Message: message,
		Data:    data,
	})
}

func Success(w http.ResponseWriter, status int, message string, data any) {
	JSON(w, status, true, message, data)
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, false, message, nil)
}
