package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Student struct
type Student struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Subjects []string `json:"subjects"`
}

// Sample Data
var students = []Student{
	{ID: 1, Name: "Alice", Subjects: []string{"Math", "Science"}},
	{ID: 2, Name: "Bob", Subjects: []string{"History", "English"}},
}

// Get all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// Get student by ID
func getStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	for _, s := range students {
		if s.ID == id {
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

// Create a new student
func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudent Student

	// Decode JSON & handle errors
	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate new ID
	newStudent.ID = len(students) + 1
	students = append(students, newStudent)
	json.NewEncoder(w).Encode(newStudent)
}

// Delete a student
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted"})
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

// Get students by subject (Query Parameter)
func getStudentsBySubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	subject := r.URL.Query().Get("subject")

	var filteredStudents []Student
	for _, s := range students {
		for _, sub := range s.Subjects {
			if sub == subject {
				filteredStudents = append(filteredStudents, s)
				break
			}
		}
	}
	json.NewEncoder(w).Encode(filteredStudents)
}

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/students", getStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getStudentByID).Methods("GET")
	router.HandleFunc("/students", createStudent).Methods("POST")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	router.HandleFunc("/students/search", getStudentsBySubject).Methods("GET")

	// Start server
	http.ListenAndServe(":8080", router)
}
