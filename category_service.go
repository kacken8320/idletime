package main

import (
	"context"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type Category struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Multiplier float64 `db:"category_multiplier"`
}

func CreateCategoryTable() {
	_, err := dbpool.Exec(context.Background(), "CREATE TABLE category (id SERIAL PRIMARY KEY,name TEXT NOT NULL,category_multiplier FLOAT NOT NULL);")
	if err != nil {
		log.Fatal(err)
	}
}

func InsertCategory(name string, multiplier float64) {
	_, err := dbpool.Exec(context.Background(), "INSERT INTO category (name, category_multiplier) VALUES ($1, $2);", name, multiplier)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCategoryMultiplier(categoryID int) float64 {
	var categoryMultiplier float64
	err := dbpool.QueryRow(context.Background(), "SELECT category_multiplier FROM category WHERE id = $1", categoryID).Scan(&categoryMultiplier)
	if err != nil {
		log.Fatal(err)
	}
	return categoryMultiplier
}

func GetAllCategories() []*Category {
	rows, err := dbpool.Query(context.Background(), "SELECT id, name, category_multiplier FROM category;")
	if err != nil {
		log.Fatal(err)
	}
	var categories []*Category
	err = pgxscan.ScanAll(&categories, rows)
	if err != nil {
		log.Fatal(err)
	}
	return categories
}
