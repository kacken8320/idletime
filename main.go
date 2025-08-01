package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

func main() {
	fmt.Println("sars")
	connStr := "host=db port=5432 user=user password=password dbname=dbname sslmode=disable"
	ctx := context.Background()

	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)

	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		err = db.Ping(ctx)
		if err == nil {
			break
		}
		log.Printf("Attempt %d: Waiting for DB... (%v)", attempts, err)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("connected to db")
}
