package models

import (
	"github.com/chuakid/govtech-onecv-2023-technical-assessment/db"
)

func SuspendStudent(email string) (err error) {
	_, err = db.DB.Exec("UPDATE students SET suspended = TRUE WHERE email = ?", email)
	return err
}
