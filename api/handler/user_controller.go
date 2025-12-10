package handler

import (
	"encoding/json"
	"net/http"
	"idletime/internal/db"
	"log"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var u db.User;
	
	// decode json
	err := json.NewDecoder(r.Body).Decode(&u);
	if err != nil {
		log.Fatal(err)
	}

	db.InsertUser(u.Username, u.Password, u.Email);

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created (vielleicht)"))
}