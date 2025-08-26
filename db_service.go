package main

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeDB(pool *pgxpool.Pool) {
	CreateCategoryTable(pool)
	CreateActivityTable(pool)
}
