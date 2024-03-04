package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mut sync.Mutex

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	resultChan := make(chan string)

	wg.Add(1)
	go crawlGoogle(&wg, resultChan)

	wg.Add(1)
	go crawlFacebook(&wg, resultChan)

	// wg.Add(1)

	usernames := []string{"user1", "user2", "user3"}

	for _, username := range usernames {
		wg.Add(1)

		go crawlTwitter(username, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println("Received: ", result)
	}

	elapsed := time.Since(start)
	// duration to seconds

	println("No of Go Routines: ", runtime.NumGoroutine())
	println("Main function took: ", elapsed.Milliseconds())
}

func crawlGoogle(wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Crawling Google")
	resultChan <- "Google"
}

func crawlFacebook(wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	time.Sleep(1 * time.Second)
	fmt.Println("Crawling Facebook")
	resultChan <- "Facebook"
}

func crawlTwitter(username string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Crawling Twitter for ", username)
	resultChan <- "Twitter for " + username

}
