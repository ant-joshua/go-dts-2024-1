package main

import "fmt"

func closureExample() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int
		for _, number := range numbers {
			if number%2 == 0 {
				result = append(result, number)
			}
		}
		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Even numbers: %d\n", evenNumbers(numbers...))
}

func closureIIFE() {
	var isPalindrome = func(str string) bool {
		var result string
		for i := len(str) - 1; i >= 0; i-- {
			result += string(str[i])
		}
		return str == result
	}("kasur rusak")

	fmt.Printf("Is palindrome: %t\n", isPalindrome)
}

func closureReturnValue(findStudent []string) func(string) bool {
	return func(name string) bool {
		for _, student := range findStudent {
			if student == name {
				return true
			}
		}
		return false
	}

}

func closureWithCallback(numbers []int, callback func(int) bool) int {
	var result int
	for _, number := range numbers {
		if callback(number) {
			result += number
		}
	}
	return result
}
