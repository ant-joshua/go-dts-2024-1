package main

func main() {

	var number = 10

	println("Before: ", number)

	changeValue(&number)

	println("After: ", number)

}

func changeValue(number *int) {
	*number = 20
}

func pointerExample() {
	var firstNumber *int
	var secondNumber *int
	var thirdNumber int

	var initialNumber = 10

	firstNumber = &initialNumber

	*firstNumber = 20

	secondNumber = firstNumber

	*secondNumber = 30

	thirdNumber = *firstNumber

	thirdNumber = 40

	println("First Number: ", *firstNumber)
	println("First Number Address: ", firstNumber)
	println("Second Number: ", *secondNumber)
	println("Second Number Address: ", secondNumber)
	println("initial Number: ", initialNumber)
	println("initial Number Address: ", &initialNumber)
	println("Third Number: ", thirdNumber)
	println("Third Number Address: ", &thirdNumber)
}
