package main

import "fmt"

func arrayExample() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var siswa = [3]string{"Budi", "Joko", "Adi"}

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", siswa)
}

func arrayModifyExample() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	numbers[0] = 10
	numbers[1] = 20

	fmt.Printf("%v\n", numbers)
}

func arrayFruitsLoop() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, v := range fruits {
		fmt.Printf("elemen %d : %s\n", i, v)
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
}

func arrayLoopMultidimenstionExample() {
	var numbers = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, number := range numbers {
		for _, n := range number {
			fmt.Println(n)
		}
	}
}
