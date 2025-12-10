package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var Dbpool *pgxpool.Pool

// Calls the functions for creating the category table and Users table.
func InitializeDB() {
	CreateCategoryTable()
	CreateUsersTable()
}
