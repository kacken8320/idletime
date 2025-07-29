package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	fmt.Println("sars")
	connStr := "host=db port=5432 user=user password=password dbname=dbname sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Attempt %d: Waiting for DB... (%v)", attempts, err)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("connected to db")
}
