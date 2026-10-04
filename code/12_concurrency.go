// Lesson 12: Concurrency in Go
// Goroutines and Channels
package main

import (
	"fmt"
	"sync"
	"time"
)

func Concurrency() {
	// GOROUTINES
	fmt.Println("=== Goroutines ===")

	// Start a goroutine
	go greetFrom("Goroutine")

	// Main continues immediately
	fmt.Println("Main function continues...")

	// Wait a bit to see goroutine output
	time.Sleep(100 * time.Millisecond)

	// Multiple goroutines
	fmt.Println("\n=== Multiple Goroutines ===")
	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Printf("Goroutine %d running\n", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	// CHANNELS
	fmt.Println("\n=== Channels ===")

	// Create a channel
	messages := make(chan string)

	// Send in goroutine
	go func() {
		messages <- "Hello from channel!"
	}()

	// Receive (blocks until message arrives)
	msg := <-messages
	fmt.Println("Received:", msg)

	// BUFFERED CHANNELS
	fmt.Println("\n=== Buffered Channels ===")

	// Channel with buffer of 2
	buffered := make(chan int, 2)

	// Can send without blocking (up to buffer size)
	buffered <- 1
	buffered <- 2
	// buffered <- 3 // Would block!

	fmt.Println("First:", <-buffered)
	fmt.Println("Second:", <-buffered)

	// CHANNEL DIRECTIONS
	fmt.Println("\n=== Channel Directions ===")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "message")
	pong(pings, pongs)

	fmt.Println("Pong:", <-pongs)

	// SELECT
	fmt.Println("\n=== Select ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "two"
	}()

	// Receive from whichever is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}

	// SELECT WITH TIMEOUT
	fmt.Println("\n=== Select with Timeout ===")

	slowChan := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		slowChan <- "slow response"
	}()

	select {
	case msg := <-slowChan:
		fmt.Println("Got:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout!")
	}

	// CLOSING CHANNELS
	fmt.Println("\n=== Closing Channels ===")

	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("All jobs received")
				done <- true
				return
			}
		}
	}()

	// Send jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	<-done

	// RANGE OVER CHANNEL
	fmt.Println("\n=== Range Over Channel ===")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println("Element:", elem)
	}

	// WAITGROUP
	fmt.Println("\n=== WaitGroup ===")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers completed")

	// MUTEX
	fmt.Println("\n=== Mutex ===")

	var counter int
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Counter:", counter)

	// WORKER POOL PATTERN
	fmt.Println("\n=== Worker Pool ===")

	numJobs := 5
	jobsChan := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobsChan, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}
	close(jobsChan)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result:", <-results)
	}
}

func greetFrom(from string) {
	fmt.Printf("Hello from %s!\n", from)
}

// Send-only channel parameter
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Receive-only and send-only channel parameters
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// Worker for worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(20 * time.Millisecond)
		results <- j * 2
	}
}
