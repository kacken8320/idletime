package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var Dbpool *pgxpool.Pool

func InitializeDB() {
	CreateCategoryTable()
	CreateActivityTable()
}
