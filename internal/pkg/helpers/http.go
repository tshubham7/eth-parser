package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Title  string `json:"title"`
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func RespondWithStatus(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
