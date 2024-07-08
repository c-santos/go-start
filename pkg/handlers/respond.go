package handlers

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, response any, code int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, err.Error(), 500)
    }
}
