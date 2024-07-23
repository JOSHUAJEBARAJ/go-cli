package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const numWorkers = 5
const numTasks = 20

func main() {
	// Create channels for tasks and results
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// Generate random tasks
	for i := 0; i < numTasks; i++ {
		tasks <- rand.Intn(100)
	}
	close(tasks)

	// Fan-out: Launch multiple worker Goroutines in the worker pool
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, results, &wg)
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()
	// Fan-in: Gather results from worker Goroutines
	var finalResults []int
	for {
		select {
		case res, ok := <-results:
			if ok {
				finalResults = append(finalResults, res)
			} else {
				results = nil
			}
		}
		if results == nil {
			break
		}
	}

	// Display results

	fmt.Println("Results:", finalResults)
}

func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Simulate task processing
		// In a real application, you would perform your actual work here
		result := task * 2

		// Send the result to the results channel
		results <- result
	}
}
