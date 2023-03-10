package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(username string, password string, db_url string, db_name string) {
	connection_string := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, db_url)
	var err error
	// create db
	DB, err = sql.Open("mysql", connection_string)
	_, err = DB.Exec("CREATE DATABASE IF NOT EXISTS " + db_name)
	_, err = DB.Exec("USE " + db_name)

	if err != nil {
		log.Fatalln("Failed to connect to db:", err)
	}
	DB.SetConnMaxIdleTime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

}
func SetupTables() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS teachers(email varchar(255) PRIMARY KEY);`)
	if err != nil {
		log.Fatalln("DB Error:", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS students(
		email varchar(255) PRIMARY KEY,
		suspended BOOLEAN DEFAULT FALSE
	);`)
	if err != nil {
		log.Fatalln("DB Error:", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS registered(
		student VARCHAR(255),
		teacher VARCHAR(255),
		FOREIGN KEY (student) REFERENCES students(email),
		FOREIGN KEY (teacher) REFERENCES teachers(email),
		CONSTRAINT PK PRIMARY KEY (student, teacher));`)
	if err != nil {
		log.Fatalln("DB Error:", err)
	}

}
