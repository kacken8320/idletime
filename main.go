package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "log"
	"os"
	"time"
	"net/http"
)

func main() {
	fmt.Println("sars")
	connStr := "host=db port=5432 user=user password=password dbname=dbname sslmode=disable"
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()

	for {
		err = dbpool.Ping(ctx)
		if err == nil {
			break
		}
		fmt.Println("Waiting for database to be available...")
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "se"+
		"lect 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	fmt.Println("connected to db")

	InitializeDB(dbpool)
	InsertCategory(dbpool, "Coroner", 4.8)
	InsertActivity(dbpool, 1, "Sarsen", 1.2, 15, 0, 0)
	fmt.Printf("Category multiplier is: %v\n", GetCategoryMultiplier(dbpool, 1))


	// healthcheck
	http.HandleFunc("/healthcheck", healthCheck) 
	http.ListenAndServe(":8320", nil) // starts an http server with a given address and handler
}
