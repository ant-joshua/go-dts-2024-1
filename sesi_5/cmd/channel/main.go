package main

import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "Hello, World!"
	c <- "Hello, Universe!"

	c <- "Hello, World!"

	//
	// c <- "Hello, World!"

	// result := <-c

	// fmt.Println(result)

	go introduce("John", c)

	go introduce("Jane", c)

	go introduce("Doe", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	defer close(c)

}

func introduce(student string, c chan string) {
	result := "Hello, " + student + "!"

	c <- result
}
