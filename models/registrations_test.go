package models

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/chuakid/govtech-onecv-2023-technical-assessment/db"
)

func TestMain(m *testing.M) {
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	db.Connect("root", "password", "127.0.0.1:3306", "test_db")
	setupTestTables(db.DB)
	m.Run()
	db.DB.Close()
}

func TestRegistration(t *testing.T) {
	err := RegisterStudentToTeacher("studentjon@gmail.com", "teacherken@gmail.com")
	if err != nil {
		t.Fatalf("Registration failed: %s", err.Error())
	}
}
func TestRegistrationWithNonExistenceTeacher(t *testing.T) {
	err := RegisterStudentToTeacher("studentjon@gmail.com", "no_exist@gmail.com")
	if err == nil {
		t.Fatalf("Registration should fail, foreign key constraints failing")
	}
}

func setupTestTables(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS teachers;")
	db.Exec("DROP TABLE IF EXISTS students;")
	db.Exec("DROP TABLE IF EXISTS registered;")

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS teachers(email varchar(255) PRIMARY KEY);`)
	if err != nil {
		log.Fatalln("db Error:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(email varchar(255) PRIMARY KEY);`)
	if err != nil {
		log.Fatalln("db Error:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS registered(
		student VARCHAR(255),
		teacher VARCHAR(255),
		FOREIGN KEY (student) REFERENCES students(email),
		FOREIGN KEY (teacher) REFERENCES teachers(email));`)
	if err != nil {
		log.Fatalln("DB Error:", err)
	}

	// Set up teachers
	_, err = db.Exec(`INSERT INTO teachers VALUES ("teacherken@gmail.com")`)
	_, err = db.Exec(`INSERT INTO teachers VALUES ("teacherjoe@gmail.com")`)

	//Set up students
	_, err = db.Exec(`INSERT INTO students VALUES ("studentjon@gmail.com")`)
	_, err = db.Exec(`INSERT INTO students VALUES ("studenthon@gmail.com")`)
	_, err = db.Exec(`INSERT INTO students VALUES ("commonstudent1@gmail.com")`)
	_, err = db.Exec(`INSERT INTO students VALUES ("commonstudent2@gmail.com")`)
	_, err = db.Exec(`INSERT INTO students VALUES ("student_only_under_teacher_ken@gmail.com")`)

	// Set up registrations
	_, err = db.Exec(`INSERT INTO registered VALUES ("student_only_under_teacher_ken@gmail.com", "teacherken@gmail.com")`)

}
