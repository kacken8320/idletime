package handler

import (
	"encoding/json"
	"idletime/internal/db"
	"log"
	"net/http"
)

// Receives an HTTP-request and forwards it to the responsible service. Returns a JSON of all categories.
//
// This function is only called from the browser or frontend (for testing maybe from curl or something).
func ControllerGetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(db.GetAllCategories())
	if err != nil {
		log.Fatal(err)
	}
}
