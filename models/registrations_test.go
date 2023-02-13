package models

import (
	"testing"
)

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
