package main

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"
)

type HealthCheckData struct {
	Message string `json:"message"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthCheckData{Message: "ok"})
}
