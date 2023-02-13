package student

import (
	"database/sql"
	"strings"

	"github.com/chuakid/govtech-onecv-2023-technical-assessment/db"
)

func SuspendStudent(email string) (err error) {
	_, err = db.DB.Exec("UPDATE students SET suspended = TRUE WHERE email = ?", email)
	return err
}

func GetCommonStudents(teachers []string) (students []string, err error) {
	query := `SELECT student 
	FROM registered
	WHERE teacher IN (?` + strings.Repeat(`,?`, len(teachers)-1) +
		`) GROUP BY student 
	HAVING count(teacher) = ?;`

	args := make([]any, len(teachers)+1)
	for i, teacher := range teachers {
		args[i] = teacher
	}
	args[len(teachers)] = len(teachers)

	var rows *sql.Rows
	rows, err = db.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var next_student string
		rows.Scan(&next_student)
		students = append(students, next_student)
	}
	return students, err
}

func GetStudentsWhoCanReceiveNotifications(teacher string, mentionedStudents [][]byte) (recipients []string, err error) {
	query := `SELECT students.email 
	FROM students
	LEFT JOIN registered ON students.email = registered.student
	WHERE
	students.suspended = FALSE AND
	(registered.teacher = ?`
	if len(mentionedStudents) == 0 {
		query += `)`
	} else {
		query += ` OR students.email IN (?` + strings.Repeat(`,?`, len(mentionedStudents)-1) + `))`
	}

	args := make([]any, len(mentionedStudents)+1)
	args[0] = teacher
	for i, student := range mentionedStudents {
		args[i+1] = string(student)[1:]
	}
	var rows *sql.Rows
	rows, err = db.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var next_student string
		rows.Scan(&next_student)
		recipients = append(recipients, next_student)
	}
	return recipients, nil
}
