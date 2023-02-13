package registration

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/chuakid/govtech-onecv-2023-technical-assessment/db"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln("Error loading .env file: ", err.Error())
	}
	MYSQL_USER := os.Getenv("TEST_MYSQL_USER")
	MYSQL_PASS := os.Getenv("TEST_MYSQL_PASS")
	DB_URL := os.Getenv("TEST_MYSQL_URL")

	const DB_NAME = "test_registrations_db"
	db.Connect(MYSQL_USER, MYSQL_PASS, DB_URL, DB_NAME)
	setupTestTables(db.DB)

	m.Run()

	// clean up
	_, err = db.DB.Exec("DROP DATABASE " + DB_NAME)
	if err != nil {
		log.Printf(err.Error())
	}
	db.DB.Close()
}

func setupTestTables(db *sql.DB) {
	_, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	_, err = db.Exec("DROP TABLE IF EXISTS teachers;")
	if err != nil {
		log.Fatalln("db Error:", err)
	}
	_, err = db.Exec("DROP TABLE IF EXISTS students;")
	if err != nil {
		log.Fatalln("db Error:", err)
	}
	_, err = db.Exec("DROP TABLE IF EXISTS registered;")
	if err != nil {
		log.Fatalln("db Error:", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS teachers(email varchar(255) PRIMARY KEY);`)
	if err != nil {
		log.Fatalln("db Error:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		email varchar(255) PRIMARY KEY,
		suspended boolean DEFAULT FALSE
	);`)
	if err != nil {
		log.Fatalln("db Error:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS registered(
		student VARCHAR(255),
		teacher VARCHAR(255),
		FOREIGN KEY (student) REFERENCES students(email),
		FOREIGN KEY (teacher) REFERENCES teachers(email),
		CONSTRAINT PK PRIMARY KEY (student, teacher));`)
	if err != nil {
		log.Fatalln("DB Error:", err)
	}

	// Set up teachers
	_, err = db.Exec(`INSERT INTO teachers VALUES ("teacherken@gmail.com")`)
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	//Set up students
	_, err = db.Exec(`INSERT INTO students VALUES ("studentjon@gmail.com", FALSE)`)
	_, err = db.Exec(`INSERT INTO students VALUES ("studenthon@gmail.com", FALSE)`)
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

}

func TestRegistration(t *testing.T) {
	students := []string{"studentjon@gmail.com"}
	err := RegisterStudentsToTeacher(students, "teacherken@gmail.com")
	if err != nil {
		t.Fatalf("Registration failed: %s", err.Error())
	}
}
func TestRegistrationWithNonExistentTeacher(t *testing.T) {
	students := []string{"studentjon@gmail.com"}
	err := RegisterStudentsToTeacher(students, "no_exist@gmail.com")
	if err == nil {
		t.Fatalf("Registration should fail, foreign key constraints failing")
	}
}
func TestRegistrationWithNonExistentStudent(t *testing.T) {
	students := []string{"no-exist@gmail.com"}
	err := RegisterStudentsToTeacher(students, "teacherken@gmail.com")
	if err == nil {
		t.Fatalf("Registration should fail, foreign key constraints failing")
	}
}
