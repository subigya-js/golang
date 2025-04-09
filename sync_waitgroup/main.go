package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d started.\n", i)
	// some taks is happening
	fmt.Printf("Worker %d ended.\n", i)

	defer wg.Done()
}

func main() {
	// fmt.Println("Explore goroutine started...")

	var wg sync.WaitGroup

	// Start 3 worker goroutines.
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the waitgroup counter
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Workers task completed.")
}
