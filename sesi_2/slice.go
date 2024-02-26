package main

import "fmt"

func sliceExample() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Printf("Jumlah element \t\t: %d\n", len(fruits))
	fmt.Printf("Kapasitas \t\t: %d\n", cap(fruits))
}

func sliceMakeExample() {
	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "grape"

	fmt.Println(fruits)

}

func sliceAppendExample() {
	var fruits = make([]string, 0)

	fruits = append(fruits, "apple", "grape", "banana")

	fmt.Printf("%v\n", fruits)
	fmt.Printf("index 0 %v\n", fruits[0])

	fmt.Printf("Jumlah element \t\t: %d\nKapasitas \t\t: %d\n", len(fruits), cap(fruits))
}

func sliceEllipsisExample() {
	var fruits = []string{"apple"}
	var fruits2 = []string{"grape", "banana", "melon"}

	// ini cara singkat nya
	fruits = append(fruits, fruits2...)

	// ini cara panjang nya
	// for _, fruit := range fruits2 {
	// 	fruits = append(fruits, fruit)
	// }

	fmt.Printf("%v", fruits)
}

func sliceCopyExample() {
	var fruits = []string{"apple", "grape", "banana"}
	var fruits2 = []string{"papaya", "orange"}

	nn := copy(fruits, fruits2)

	fmt.Println(fruits)
	fmt.Println(fruits2)
	fmt.Println(nn)
}

func sliceSlicingExample() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:4]
	var bFruits = fruits[1:4]

	fmt.Println(aFruits)
	fmt.Println(bFruits)
}
