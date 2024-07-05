package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func baseHandler(w http.ResponseWriter, r *http.Request, response Response) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, you've called me!",
		Status:  200,
	}

	baseHandler(w, r, response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// response := Response{
	// 	Message: "Hello, %s\n", r.PathValue("name"),
	// 	Status: 200,
	// }

	// baseHandler(w, r, response)
}

