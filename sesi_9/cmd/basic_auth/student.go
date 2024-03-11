package main

import "net/http"

var students = []*Student{}

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}

	OutputJsonEncode(w, GetStudents())
}

func GetStudents() []*Student {
	return students
}

func init() {
	students = append(students, &Student{ID: 1, Name: "wick", Grade: 2})
	students = append(students, &Student{ID: 2, Name: "john", Grade: 3})
	students = append(students, &Student{ID: 3, Name: "doe", Grade: 3})
}

func SelectStudentByID(id int) *Student {
	for _, each := range students {
		if each.ID == id {
			return each
		}
	}
	return nil
}
