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
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASS := os.Getenv("MYSQL_PASS")
	DB_URL := os.Getenv("MYSQL_URL")

	db.Connect(MYSQL_USER, MYSQL_PASS, DB_URL, "technical_assessment")
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
	log.Printf("Server up on %v", PORT)
	http.ListenAndServe(":"+PORT, r)
}
