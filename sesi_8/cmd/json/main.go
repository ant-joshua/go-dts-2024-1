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

	var jsonString = `{"name": "John Doe", "email": "antoniusjoshua47@gmail.com", "age": 27}`

	var result Employee

	// Unmarshal JSON string to struct

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		println(err.Error())
		return
	}

	var resultMap map[string]any

	// println("Name  :", result.Name)
	// println("Email :", result.Email)
	// println("Age   :", result.Age)

	err = json.Unmarshal([]byte(jsonString), &resultMap)

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println("Name  :", resultMap["name"])

	jsonData, err := json.Marshal(result)

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(jsonData))

	jsonDataMap, err := json.Marshal(resultMap)

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(jsonDataMap))

	var temp any

	err = json.Unmarshal(jsonData, &temp)

	if err != nil {
		println(err.Error())
		return
	}

	var resultTemp = temp.(map[string]interface{})

	fmt.Println("resultTemp", resultTemp)

	// println("Name  :", resultMap["name"])
	// println("Email :", resultMap["email"])
	// println("Age   :", resultMap["age"])

}
