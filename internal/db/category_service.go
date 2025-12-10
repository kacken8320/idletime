package db

import (
	"context"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type Category struct {
	ID                int     `db:"id"`
	UserID            int     `db:"user_id"`
	ParentCategoryID  int     `db:"parent_category_id"` //=0 if no parent exists
	Name              string  `db:"name"`
	Multiplier        float64 `db:"multiplier"`
	TimeSpentInMins   int     `db:"time_spent_in_mins"`
	MinimalTimeInMins int     `db:"minimal_time_in_mins"` //=0 if not an activity
	SkipCounter       int     `db:"skip_counter"`         //=0 if not an activity
	IsActivity        bool    `db:"is_activity"`
}

func CreateCategoryTable() {
	_, err := Dbpool.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS category (id SERIAL PRIMARY KEY, user_id INTEGER NOT NULL, parent_category_id INTEGER NOT NULL, name TEXT NOT NULL, multiplier FLOAT NOT NULL, time_spent_in_mins INTEGER NOT NULL, minimal_time_in_mins INTEGER NOT NULL, skip_counter INTEGER NOT NULL, is_activity BOOLEAN NOT NULL);")
	if err != nil {
		log.Fatal(err)
	}
}

func InsertCategory(userID int, parentCategoryID int, name string, multiplier float64, timeSpentInMins int, minimalTimeInMins int, isActivity bool) {
	_, err := Dbpool.Exec(context.Background(), "INSERT INTO category (user_id, parent_category_id, name, multiplier, time_spent_in_mins, minimal_time_in_mins, skip_counter, is_activity) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);", userID, parentCategoryID, name, multiplier, 0, timeSpentInMins, minimalTimeInMins, isActivity)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMultiplier(categoryID int) float64 {
	var Multiplier float64
	err := Dbpool.QueryRow(context.Background(), "SELECT multiplier FROM category WHERE id = $1", categoryID).Scan(&Multiplier)
	if err != nil {
		log.Fatal(err)
	}
	return Multiplier
}

func GetAllCategories() []Category {
	rows, err := Dbpool.Query(context.Background(), "SELECT id, user_id, parent_category_id, name, multiplier, time_spent_in_mins, minimal_time_in_mins, skip_counter, is_activity FROM category;")
	if err != nil {
		log.Fatal(err)
	}
	var categories []Category
	err = pgxscan.ScanAll(&categories, rows)
	if err != nil {
		log.Fatal(err)
	}
	return categories
}
