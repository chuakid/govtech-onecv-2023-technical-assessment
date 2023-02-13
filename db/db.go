package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(username string, password string, db_url string) {
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/main", username, password, db_url))
	if err != nil {
		log.Fatalln("Failed to connect to db:", err)
	}
	DB.SetConnMaxIdleTime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
