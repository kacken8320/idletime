package db

import (
	"context"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func CreateUsersTable() {
	_, err := Dbpool.Exec(context.Background(), "CREATE EXTENSION IF NOT EXISTS citext;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = Dbpool.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username TEXT NOT NULL, password TEXT NOT NULL, email CITEXT NOT NULL UNIQUE);")
	if err != nil {
		log.Fatal(err)
	}
}

func InsertUser(username string, password string, email string) {
	_, err := Dbpool.Exec(context.Background(),
		"INSERT INTO users (username, password, email) VALUES ($1, $2, $3);", username, password, email)
	if err != nil {
		log.Println(err)
	}
}

func GetAllUsers() []User {
	rows, err := Dbpool.Query(context.Background(), "SELECT id, username, password, email FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	err = pgxscan.ScanAll(&users, rows)
	if err != nil {
		log.Fatal(err)
	}
	return users
}
