package main

// package advanced

import (
	"fmt"
	"time"
)

// RunConcurrencyExample demonstrates goroutines and channels
func RunConcurrencyExample() {
	fmt.Println("7️⃣  Concurrency (Goroutines & Channels)")

	// Simple goroutine
	fmt.Println("   Starting goroutines...")
	go printNumbers("A")
	go printNumbers("B")

	// Wait a bit for goroutines to execute
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// Channels - for communication between goroutines
	fmt.Println("   Using channels:")
	messages := make(chan string)

	// Send message in goroutine
	go func() {
		messages <- "Hello from goroutine!"
	}()

	// Receive message
	msg := <-messages
	fmt.Printf("   Received: %s\n", msg)

	// Buffered channels
	fmt.Println("   Buffered channel:")
	numbers := make(chan int, 3)
	numbers <- 1
	numbers <- 2
	numbers <- 3
	close(numbers) // Close channel when done sending
	//!To Understand throughly
	// Receive all values
	for num := range numbers {
		fmt.Printf("   Got: %d\n", num)
	}

	// Channel with worker goroutines
	fmt.Println("   Worker pool example:")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}

	// Select statement - choose from multiple channels
	fmt.Println("   Select statement:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(5 * time.Millisecond)
		ch2 <- "Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("   Received from %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("   Received from %s\n", msg2)
		}
	}
}

func printNumbers(prefix string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("   [%s-%d] ", prefix, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("   Worker %d processing job %d\n", id, j)
		time.Sleep(10 * time.Millisecond)
		results <- j * 2
	}
}

func main() {
	RunConcurrencyExample()
}
