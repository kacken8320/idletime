package db

import (
	"context"
	"log"
)

func CreateActivityTable() {
	_, err := Dbpool.Exec(context.Background(), "CREATE TABLE activity (id SERIAL PRIMARY KEY,category_id INTEGER NOT NULL,name TEXT NOT NULL, activity_multiplier FLOAT NOT NULL,minimal_time_in_minutes INTEGER NOT NULL,time_spent_in_minutes INTEGER NOT NULL,skip_counter INTEGER NOT NULL);")
	if err != nil {
		log.Fatal(err)
	}
}

func InsertActivity(categoryID int, name string, multiplier float64, minTime int, timeSpent int, skipCount int) {
	_, err := Dbpool.Exec(context.Background(), "INSERT INTO activity (category_id, name, activity_multiplier, minimal_time_in_minutes, time_spent_in_minutes, skip_counter) VALUES ($1, $2, $3, $4, $5, $6);", categoryID, name, multiplier, minTime, timeSpent, skipCount)
	if err != nil {
		log.Fatal(err)
	}
}
