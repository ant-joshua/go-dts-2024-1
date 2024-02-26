package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	println(greetMultiNames("Hello", []string{"John", "Doe", "Jane", "Doe"}))

	println(greetMultiVariadic("Hello", "John", "Doe", "Jane", "Doe", "joshua", "antonius", "joseph"))

	area, circumference := calculateDiameter(15)

	fmt.Printf("Area: %2.f \n", area)
	fmt.Printf("Circumference: %2.f ", circumference)

	studentScore("Antonius Joshua", 70, 80, 90, 100)

	// _, notes := studentNotes("Antonius Joshua", "A", "B", "C", "D", 100)

	// fmt.Printf("Notes %s", notes)

	// closureExample()

	// closureIIFE()

	// var findOddNumbers = closureWithCallback([]int{1, 2, 3, 4, 5}, isOdd)

	// fmt.Printf("Odd numbers: %d\n", findOddNumbers)
}

func isOdd(m int) bool {
	return m%2 != 0
}

func greet(firstName, lastName string) string {

	return "Hello" + " " + firstName + " " + lastName
}

func greetMultiNames(message string, names []string) string {
	var stringJoins = strings.Join(names, ", ")

	return message + " " + stringJoins
}

func greetMultiVariadic(message string, names ...string) string {
	var stringJoins = strings.Join(names, ", ")

	return message + " " + stringJoins
}

func greetMultiReturn(firstName, lastName, address string) (string, string, string) {

	return "Hello", firstName, lastName
}

func calculateDiameter(d float64) (float64, float64) {
	const pi = 3.14

	var area = math.Pi * math.Pow(d/2, 2)

	var circumference = math.Pi * d

	return area, circumference
}

func calculateDiameterPredefined() (area float64, circumference float64) {

	const pi = 3.14
	var d = 15.0

	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

func studentScore(name string, scores ...int) (string, int) {
	var totalScore int

	for _, score := range scores {
		totalScore += score
	}

	return name, totalScore
}

func studentNotes(name string, notes ...any) (string, string) {

	var stringJoins string

	for _, note := range notes {
		stringJoins += note.(string) + " "
	}

	return name, stringJoins
}
