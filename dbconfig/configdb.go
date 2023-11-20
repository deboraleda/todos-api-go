package dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func ConfigDB() {
	connStr := "postgresql://deboraleda:C3Td6cDGLFXr@ep-flat-darkness-02734092.us-east-2.aws.neon.tech/todos"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
}
