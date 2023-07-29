## Concurrency patterns

Concurrency pattern such as worker pools provide a way to limit the number of concurrent operations that are performed to prevent resource exhaustion and manage contention.

### Worker pools

1. We have a list of tasks to be performed (jobs)
2. No of workers are limited (worker pool)
3. Results are collected (results)


Why  jobs is not equal to workers?

- We have network bandwith limitations