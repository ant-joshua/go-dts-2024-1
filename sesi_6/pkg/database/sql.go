package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "hacktiv8_dts_kominfo"
)

var (
	DB  *sql.DB
	err error
)

func SQLInit() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	return DB
}
