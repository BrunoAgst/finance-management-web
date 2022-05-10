package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConect() *sql.DB {
	connection := "user=admin dbname=postgres password=layla123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
