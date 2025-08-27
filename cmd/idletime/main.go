package main

import (
	"context"
	"fmt"
	handler2 "idletime/api/handler"
	"idletime/internal/db"
	_ "log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("sars")
	connStr := "host=db port=5432 user=user password=password dbname=dbname sslmode=disable"
	ctx := context.Background()

	var err error
	db.Dbpool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer db.Dbpool.Close()

	for {
		err = db.Dbpool.Ping(ctx)
		if err == nil {
			break
		}
		fmt.Println("Waiting for database to be available...")
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	var greeting string
	err = db.Dbpool.QueryRow(context.Background(), "se"+
		"lect 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	fmt.Println("connected to db")
	db.InitializeDB()
	db.InsertCategory("Coroner", 4.8)
	db.InsertCategory("Sers", 9.6)
	db.InsertActivity(1, "Sarsen", 1.2, 15, 0, 0)
	fmt.Printf("Category multiplier is: %v\n", db.GetCategoryMultiplier(1))
	categories := db.GetAllCategories()
	for _, u := range categories {
		fmt.Println(u.ID, u.Name, u.Multiplier)
	}

	// healthcheck
	http.HandleFunc("/healthcheck", handler2.HealthCheck)
	http.HandleFunc("/GetAllCategories", handler2.ControllerGetAllCategories)
	http.ListenAndServe(":8320", nil) // starts an handler server with a given address and handler
}
