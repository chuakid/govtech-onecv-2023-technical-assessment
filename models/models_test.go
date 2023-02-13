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
	_, err = db.Exec(`INSERT INTO teachers VALUES ("teacherjoe@gmail.com")`)

	//Set up students
	_, err = db.Exec(`INSERT INTO students VALUES ("studentjon@gmail.com", FALSE)`)
	_, err = db.Exec(`INSERT INTO students VALUES ("studenthon@gmail.com", FALSE)`)
	_, err = db.Exec(`INSERT INTO students VALUES ("commonstudent1@gmail.com", FALSE)`)
	_, err = db.Exec(`INSERT INTO students VALUES ("commonstudent2@gmail.com", FALSE)`)
	_, err = db.Exec(`INSERT INTO students VALUES ("student_only_under_teacher_ken@gmail.com", FALSE)`)
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
	// Set up registrations
	_, err = db.Exec(`INSERT INTO registered VALUES ("student_only_under_teacher_ken@gmail.com", "teacherken@gmail.com")`)
	_, err = db.Exec(`INSERT INTO registered VALUES ("commonstudent1@gmail.com", "teacherken@gmail.com")`)
	_, err = db.Exec(`INSERT INTO registered VALUES ("commonstudent2@gmail.com", "teacherken@gmail.com")`)
	_, err = db.Exec(`INSERT INTO registered VALUES ("commonstudent1@gmail.com", "teacherjoe@gmail.com")`)
	_, err = db.Exec(`INSERT INTO registered VALUES ("commonstudent2@gmail.com", "teacherjoe@gmail.com")`)
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

}
