package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chuakid/govtech-onecv-2023-technical-assessment/models"
)

func registerStudents(w http.ResponseWriter, r *http.Request) {
	log.Printf("Register Students endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	type DataFormat struct {
		Teacher  string
		Students []string
	}
	var studentsAndTeachers DataFormat
	err := json.NewDecoder(r.Body).Decode(&studentsAndTeachers)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Formatting error",
		})
		return
	}
	for _, student := range studentsAndTeachers.Students {
		err = models.RegisterStudentToTeacher(student, studentsAndTeachers.Teacher)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Either student or teacher not found",
			})
			return
		}
	}
	w.WriteHeader(204)
}
