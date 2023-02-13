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
	err = models.RegisterStudentsToTeacher(studentsAndTeachers.Students, studentsAndTeachers.Teacher)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Either student or teacher not found",
		})
		return
	}

	w.WriteHeader(204)
}
func suspendStudent(w http.ResponseWriter, r *http.Request) {
	log.Printf("Suspend student endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	type DataFormat struct {
		Student string
	}
	var student DataFormat
	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Formatting error",
		})
		return
	}
	err = models.SuspendStudent(student.Student)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Suspend fail",
		})
		log.Printf(err.Error())
		return
	}

	w.WriteHeader(204)
}

func getCommonStudents(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get common students endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	teachers := r.URL.Query()["teacher"]
	if len(teachers) == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "No teachers in query",
		})
		return
	}
	res, err := models.GetCommonStudents(teachers)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Failed to get common students",
		})
		log.Printf(err.Error())
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"students": res,
	})

}
