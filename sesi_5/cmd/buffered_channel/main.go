package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c1 := make(chan int, 2)

	go func(c chan int) {
		for i := 0; i < 10; i++ {
			fmt.Printf("func goroutine: #%d start sending data to channel", i)
			c <- i
			fmt.Printf("func goroutine: #%d data sent to channel", i)
		}

		close(c)
	}(c1)

	fmt.Println("main goroutine: reading data from channel")
	time.Sleep(2 * time.Second)

	for v := range c1 {
		fmt.Println("main goroutine: received data from channel:", v)
	}

	// os.Exit()
	// var err error

	_, err := strconv.Atoi("10aaasdasd")

	if err != nil {
		panic(err)
	}

}
