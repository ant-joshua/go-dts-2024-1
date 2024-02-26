package main

import "fmt"

func stringByteExample() {
	var name string = "Budi"

	for i := 0; i < len(name); i++ {
		fmt.Printf("%d : %d\n", i, name[i])
	}
}

func runeByteExample() {
	str := "makan"

	str2 := "mânca"

	fmt.Printf("len(str) = %d\n", len(str))
	fmt.Printf("len(str2) = %d\n", len(str2))
}
