package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	fmt.Println("Server is running on port 5000")

	http.HandleFunc("/employees", getEmployeeListHtml)
	http.HandleFunc("/api/employees", getEmployeeList)
	http.HandleFunc("/employees/create", createEmployee)
	http.HandleFunc("/employees/find", getEmployee)

	http.ListenAndServe(":5000", nil)
}

type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

var employees = []Employee{
	{ID: "123e4567-e89b-12d3-a456-426614174000", Name: "John Doe", Age: 25, Division: "Engineering"},
	{ID: "123e4567-e89b-12d3-a456-426614174001", Name: "Jane Doe", Age: 30, Division: "Marketing"},
}

func getEmployeeList(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Set Header to return JSON
	w.Header().Set("Content-Type", "application/json")
	// Set Status to 200 (OK)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func getEmployeeListHtml(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Set Header to return JSON

	// Set Status to 200 (OK)

	tpl, err := template.ParseFiles("./template.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	tpl.ExecuteTemplate(w, "employees", employees)
	return
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Set Header to return JSON
	w.Header().Set("Content-Type", "application/json")
	// Set Status to 201 (Created)
	w.WriteHeader(http.StatusCreated)

	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)

	newUUID := uuid.New()

	employee.ID = newUUID.String()

	employees = append(employees, employee)

	json.NewEncoder(w).Encode(employee)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/employees/find/")

	fmt.Println(id)

	if id == "" {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}

	_, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)

		return
	}

	// id := r.URL.Query().Get("id")

	var employeeData Employee

	for _, employee := range employees {
		if employee.ID == id {
			employeeData = employee
		}
	}

	// Set Header to return JSON
	w.Header().Set("Content-Type", "application/json")
	// Set Status to 200 (OK)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeeData)
}
