package main

import (
	"fmt"
	"time"
)

/** code explanation
First going to spin three go routines
Next going to load the data in the jobs channel (input)
Next we are going to close the job channel, This is necessary for the for run loop
Next the worker will start execute
**/
func main() {
	const numJobs = 10
	jobsChan := make(chan int, numJobs)
	completedJobsChan := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobsChan, completedJobsChan)
	}

	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}
	close(jobsChan)
	for a := 1; a <= numJobs; a++ {
		<-completedJobsChan
	}

}

func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) { // single direction channeled, send to completed chan and recieve from the jobschan

	for j := range jobsChan {
		fmt.Println("worker", id, "started", j, len(jobsChan), "Remaining")
		time.Sleep(time.Second * 2)
		fmt.Println("worker", id, "finised job", j)
		completedJobsChan <- j
	}
}
