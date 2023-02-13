package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chuakid/govtech-onecv-2023-technical-assessment/db"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Set up logger
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_pass := os.Getenv("MYSQL_PASS")
	db_url := os.Getenv("MYSQL_URL")

	db.Connect(mysql_user, mysql_pass, db_url, "technical_assessment")
	log.Println("DB connected")
	defer db.DB.Close()

	// Create tables
	db.SetupTables()
	log.Println("Tables setup")

	// Router
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/register", registerStudents)
		r.Post("/suspend", suspendStudent)
		r.Post("/retrievefornotifications", getForNotifications)
		r.Get("/commonstudents", getCommonStudents)
	})
	http.ListenAndServe(":8000", r)

}
