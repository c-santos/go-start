package handlers

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, response any) {
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, err.Error(), 500)
    }
}
