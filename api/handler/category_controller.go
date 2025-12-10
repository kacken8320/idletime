package handler

import (
	"encoding/json"
	"idletime/internal/db"
	"log"
	"net/http"
)

func ControllerGetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(db.GetAllCategories())
	if err != nil {
		log.Fatal(err)
	}
}
