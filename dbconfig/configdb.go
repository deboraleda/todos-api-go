package dbconfig

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnvFile() error {
	return godotenv.Load()
}

var DB *sql.DB
var err error

func ConfigDB() {
	loadEnvFile()
	databaseURL := os.Getenv("DATABASE_URL")
	connStr := databaseURL
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
}
