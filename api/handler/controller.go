package handler

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type HealthCheckData struct {
	Message string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthCheckData{Message: "ok"})
	fmt.Println("health checked")
}
