## Concurrency patterns

Concurrency pattern such as worker pools provide a way to limit the number of concurrent operations that are performed to prevent resource exhaustion and manage contention.

### Worker pools

1. We have a list of tasks to be performed (jobs)
2. No of workers are limited (worker pool)
3. Results are collected (results)


Why  jobs is not equal to workers?

- We have network bandwith limitations


## Worker pool vs fan in and fan out 

You can use a Worker Pool to manage a fixed number of worker Goroutines that process tasks concurrently.
Then, you can use the Fan-out pattern to distribute tasks from a central Goroutine to the worker Goroutines in the pool.
Finally, you can use the Fan-in pattern to aggregate the results produced by the worker Goroutines into a single channel for further processing or display.