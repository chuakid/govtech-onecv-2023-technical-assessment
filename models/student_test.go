package models

import "testing"

func TestSuspendStudent(t *testing.T) {
	err := SuspendStudent("studentjon@gmail.com")
	if err != nil {
		t.Fatalf("Suspension fail, %s", err)
	}
}

func TestGetCommonStudents(t *testing.T) {
	students, err := GetCommonStudents([]string{"teacherken@gmail.com"})
	if err != nil {
		t.Fatalf("Get common students fail, %s", err)
	}
	if len(students) != 4 {
		t.Fatalf("Get common students fail, students: %v, expected: %v", students, `[
			commonstudent1@gmail.com,
			commonstudent2@gmail.com,
			studentjon@gmail.com,
			student_only_under_teacher_ken@gmail.com@gmail.com
		]`)
	}
}

func TestGetCommonStudentsWithTwoTeachers(t *testing.T) {
	students, err := GetCommonStudents([]string{"teacherken@gmail.com", "teacherjoe@gmail.com"})
	if err != nil {
		t.Fatalf("Get common students fail, %s", err)
	}
	if len(students) != 2 {
		t.Fatalf("Get common students fail, students: %v, expected: %v", students, `[
			commonstudent1@gmail.com,
			commonstudent2@gmail.com,
		]`)
	}
}
