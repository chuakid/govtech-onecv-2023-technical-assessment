package models

import "github.com/chuakid/govtech-onecv-2023-technical-assessment/db"

func RegisterStudentToTeacher(student string, teacher string) (err error) {
	_, err = db.DB.Exec("INSERT INTO registered VALUES (?,?)", student, teacher)
	return err
}
