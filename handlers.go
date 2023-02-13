package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/chuakid/govtech-onecv-2023-technical-assessment/models/registration"
	"github.com/chuakid/govtech-onecv-2023-technical-assessment/models/student"
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
	err = registration.RegisterStudentsToTeacher(studentsAndTeachers.Students, studentsAndTeachers.Teacher)
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
	var studentToBeSuspended DataFormat
	err := json.NewDecoder(r.Body).Decode(&studentToBeSuspended)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Formatting error",
		})
		return
	}
	err = student.SuspendStudent(studentToBeSuspended.Student)
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
	res, err := student.GetCommonStudents(teachers)
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
func getForNotifications(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get for notifications endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	type DataFormat struct {
		Teacher      string
		Notification string
	}
	var notificationAndTeacher DataFormat
	err := json.NewDecoder(r.Body).Decode(&notificationAndTeacher)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Formatting error",
		})
		return
	}

	regex, err := regexp.Compile(`@\w*@\w*\.\w*`)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	mentioned := regex.FindAll([]byte(notificationAndTeacher.Notification), -1)
	recipients, err := student.GetStudentsWhoCanReceiveNotifications(notificationAndTeacher.Teacher, mentioned)
	if err != nil {
		log.Printf("Error getting students who can receive notifications %s", err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Error retrieving data",
		})
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"recipients": recipients,
	})
	return

}
