package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. SIMPLE GOROUTINE
func sayHello(name string) {
	fmt.Printf("   Hello, %s! (from a goroutine)\n", name)
}

// 2. CHANNELS
// Channels are the pipes that connect concurrent goroutines.
func sendData(ch chan string) {
	fmt.Println("   Sending data to channel...")
	time.Sleep(1 * time.Second) // Simulate work
	ch <- "Message from the channel! 🚀"
}

// 3. WORKER POOL EXAMPLE
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("   Worker %d started job %d\n", id, j)
		time.Sleep(500 * time.Millisecond) // Simulate work
		fmt.Printf("   Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	fmt.Println("--- Concurrency in Go ---")

	// 1. GOROUTINES
	// Use 'go' keyword to starch a new thread of execution.
	fmt.Println("\n1. Launching Goroutines:")
	go sayHello("Alice")
	go sayHello("Bob")
	
	// We need to wait a bit because main might exit before goroutines finish.
	time.Sleep(100 * time.Millisecond)

	// 2. CHANNELS
	fmt.Println("\n2. Using Channels:")
	msgChannel := make(chan string) // Unbuffered channel
	go sendData(msgChannel)

	// Receiving from a channel is a blocking operation.
	msg := <-msgChannel
	fmt.Println("   Received:", msg)

	// 3. WAITGROUPS
	// Safe way to wait for multiple goroutines to finish without 'time.Sleep'.
	fmt.Println("\n3. Synchronization with WaitGroups:")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("   Task %d is running...\n", id)
			time.Sleep(200 * time.Millisecond)
		}(i)
	}
	wg.Wait() // Blocks until all three Done() calls are made
	fmt.Println("   All tasks finished!")

	// 4. SELECT STATEMENT
	// Allows a goroutine to wait on multiple communication operations.
	fmt.Println("\n4. Select Statement (Handling multiple channels):")
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		chan1 <- "From Chan 1"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		chan2 <- "From Chan 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("   Received:", msg1)
		case msg2 := <-chan2:
			fmt.Println("   Received:", msg2)
		}
	}

	// 5. WORKER POOL PATTERN
	fmt.Println("\n5. Worker Pool Pattern:")
	const numJobs = 5
	jobs := make(chan int, numJobs)    // Buffered channel
	results := make(chan int, numJobs)
	
	var poolWg sync.WaitGroup
	
	// Start 3 workers
	for w := 1; w <= 3; w++ {
		poolWg.Add(1)
		go worker(w, jobs, results, &poolWg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Important to close channels when no more data will be sent

	// Wait for all workers to finish
	poolWg.Wait()
	close(results)

	// Collect results
	fmt.Print("   Final Results: ")
	for res := range results {
		fmt.Printf("%d ", res)
	}
	fmt.Println()

	/*
	SUMMARY:
	- 'go' keyword starts a concurrent thread (goroutine).
	- Use channels (chan) to share data safely between goroutines.
	- Don't communicate by sharing memory; share memory by communicating.
	- WaitGroups (sync.WaitGroup) are essential for waiting on multiple tasks.
	- 'select' is like a switch statement for channels.
	- Close channels when they are no longer needed to prevent deadlocks or leaks.
	*/
}
