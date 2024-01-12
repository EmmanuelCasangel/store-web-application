package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	conection := "user=postgres dbname=postgres host=localhost password=123456 sslmode=disable"

	db, err := sql.Open("postgres", conection)

	if err != nil {
		panic(err)
	}

	return db

}
