package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ErrorMessage *string      `json:"errorMessage,omitempty"`
	Data         *interface{} `json:"Data,omitempty"`
}

func GenerateResponse(w http.ResponseWriter, data interface{}, status int) {
	resp := Response{}
	if data != nil {
		resp = Response{
			Data: &data,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func GenerateErrorResponse(w http.ResponseWriter, message string, status int) {
	resp := Response{
		ErrorMessage: &message,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
