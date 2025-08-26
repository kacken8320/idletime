package main

import (
	"fmt"
	"net/http"
	_ "encoding/json"
)

type HealthCheckData struct {
	Message	string `json:"message"`
}

		// return object	data about incoming request
func healthCheck(w http.ResponseWriter, r *http.Request) {
	// response code
	w.WriteHeader(http.StatusOK)

	// response body
	fmt.Fprintln(w, "up")
}

