package main

import (
	"log"
	"os"

	"github.com/chuakid/govtech-onecv-2023-technical-assessment/db"
)

func main() {
	// Set up logger
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_pass := os.Getenv("MYSQL_PASS")
	db_url := os.Getenv("MYSQL_URL")

	db.Connect(mysql_user, mysql_pass, db_url)
	log.Println("DB connected")
}
