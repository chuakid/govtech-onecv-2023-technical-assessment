package registration

import "github.com/chuakid/govtech-onecv-2023-technical-assessment/db"

func RegisterStudentsToTeacher(students []string, teacher string) (err error) {
	tx, err := db.DB.Begin()
	for _, student := range students {
		_, err = db.DB.Exec("INSERT INTO registered VALUES (?,?)", student, teacher)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return err
}
