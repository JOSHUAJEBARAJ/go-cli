package main

import (
	"fmt"
	"math/rand"
)

const numWorkers = 5
const numTasks = 20

func main() {
	// Create channels for tasks and results
	taskQueue := make(chan int, numTasks)
	resultQueue := make(chan int, numTasks)
	done := make(chan struct{}) // Channel to signal when all workers are done

	// Generate random tasks
	for i := 0; i < numTasks; i++ {
		taskQueue <- rand.Intn(100)
	}
	close(taskQueue)

	// Fan-out: Launch multiple worker Goroutines in the worker pool
	for i := 0; i < numWorkers; i++ {
		go worker(taskQueue, resultQueue, done)
	}

	// Wait for all workers to finish
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	close(resultQueue)

	// Fan-in: Gather results from worker Goroutines
	var finalResults []int
	for res := range resultQueue {
		finalResults = append(finalResults, res)
	}

	// Display results
	fmt.Println("Generated tasks:", numTasks)
	fmt.Println("Processed tasks:", len(finalResults))
	fmt.Println("Results:", finalResults)
}

func worker(taskQueue <-chan int, resultQueue chan<- int, done chan<- struct{}) {
	for task := range taskQueue {
		// Simulate task processing
		// In a real application, you would perform your actual work here
		result := task * 2

		// Send the result to the results channel
		resultQueue <- result
	}
	done <- struct{}{} // Signal that this worker is done
}
