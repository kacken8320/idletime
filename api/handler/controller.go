package handler

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// The string contains a message wether the server is up. It's in JSON-format.
type HealthCheckData struct {
	Message string `json:"message"`
}

// Checks the health :3
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthCheckData{Message: "ok"})
	fmt.Println("health checked")
}
