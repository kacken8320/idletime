package main

import (
	"encoding/json"
	_ "encoding/json"
	"log"
	"net/http"
)

func ControllerGetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(GetAllCategories())
	if err != nil {
		log.Fatal(err)
	}
}
