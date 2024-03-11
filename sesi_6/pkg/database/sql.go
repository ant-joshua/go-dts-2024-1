package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "hacktiv8_dts_kominfo"
// )

var (
	DB  *sql.DB
	err error
)

func SQLInit() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", viper.GetString("DB_HOST"), viper.GetInt("DB_PORT"), viper.GetString("DB_USER"), viper.GetString("DB_PASS"), viper.GetString("DB_NAME"))

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
