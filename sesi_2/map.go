package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Hobbies []string
} // class Person

// Record<string,any> = object

func mapExample() {
	var person map[string]any

	var peopleList []map[string]any

	person = map[string]any{
		"name":    "Budi",
		"age":     23,
		"hobbies": []string{"eating", "sleeping"},
	}

	for key, value := range person {
		fmt.Printf("%s : %v\n", key, value)
	}

	delete(person, "age")

	value, isExists := person["age"]

	if !isExists {
		fmt.Println("data age tidak ada")
	}

	peopleList = append(peopleList, person)

	fmt.Println(value)

	fmt.Println(isExists)

	fmt.Println(person)

}
