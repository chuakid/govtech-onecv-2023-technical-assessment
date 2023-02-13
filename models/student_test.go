package models

import "testing"

func TestSuspendStudent(t *testing.T) {
	err := SuspendStudent("studentjon@gmail.com")
	if err != nil {
		t.Fatalf("Suspension fail, %s", err)
	}
}
