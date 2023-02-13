package registration

import "github.com/chuakid/govtech-onecv-2023-technical-assessment/db"

func RegisterStudentsToTeacher(students []string, teacher string) (err error) {
	for _, student := range students {
		_, err = db.DB.Exec("INSERT INTO registered VALUES (?,?)", student, teacher)
		if err != nil {
			return err
		}
	}
	return err
}
