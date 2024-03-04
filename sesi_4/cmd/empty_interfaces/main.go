package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var random any

	random = 1

	random = "random"

	random = struct{}{}

	random = []int{1, 2, 3}

	random = "20"

	fmt.Printf("Type of random: %T\n", random)
	fmt.Printf("Value of random: %v\n", random)

	if value, ok := random.(int); !ok {
		fmt.Println("Value is not an int")
	} else {
		fmt.Println("Value is an int:", value)
	}

	fmt.Println("type of with reflect", reflect.TypeOf(random))
	fmt.Println("value of with reflect", reflect.ValueOf(random))

	if reflect.TypeOf(random).Kind() == reflect.String {

		convertToInt, _ := random.(string)

		convertedType, err := strconv.Atoi(convertToInt)
		if err != nil {
			fmt.Println("Error converting to int")
		}

		fmt.Println("Value is a string:", convertedType)
	} else if reflect.TypeOf(random).Kind() == reflect.Slice {
		fmt.Println("Value is a slice")
	} else if reflect.TypeOf(random).Kind() == reflect.Struct {
		fmt.Println("Value is a struct")
	} else if reflect.TypeOf(random).Kind() == reflect.Int {
		fmt.Println("Value is an int")

		fmt.Println("Value is an int:", random.(int))

	} else {
		fmt.Println("Value is not a string, slice or struct")
	}

	// json.Marshal()

	// rs := []any{1, "random", struct{}{}, []int{1, 2, 3}, "random"}

	// for _, r := range rs {
	// 	fmt.Printf("Type of r: %T\n", r)
	// 	fmt.Printf("Value of r: %v\n", r)
	// }

}
