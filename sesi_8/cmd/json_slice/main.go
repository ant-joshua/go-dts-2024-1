package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	var jsonString = `[
		{"name": "John Doe", "email": "antoniusjoshua47@gmail.com", "age": 27},
		{"name": "John Doe", "email": "antoniusjoshua47@gmail.com", "age": 27}
	]`

	var result []Employee

	err := json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Println("Employee", i)
		fmt.Println("Name  :", v.Name)
		fmt.Println("Email :", v.Email)
		fmt.Println("Age   :", v.Age)
	}
}
