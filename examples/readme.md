## Context in golang

https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

- Context allows to cancellation propagation , Let's try this to understand with the example, lets say you run a shop and someone orders a food and you send your employee to buy the ingredients and the customer cancels the order, now you have to cancel the order of the ingredients as well, so you have to propagate the cancellation to the employee as well, this is what context does, it propagates the cancellation to the goroutines as well.


## Dummy context 

- Use context.TODO function is the one way to create an empty context
- Use context.Background function is the other way to create an empty context


## Using Data withing context
- One of the benefits of using context is that it allows you to pass data between functions without having to pass the data as an argument to each function

- context provides a way to store data within the context and retrieve it when needed

- context.WithValue function allows you to store data within the context and retrieve it when needed

- When using contexts, it’s important to know that the values stored in a specific context.Context are immutable
## Ending a context 

- Another powerful tool context.context provides is a way to signal to any functions using it that context has ended and should be considered complete'
- context provides Done method which checks if the context has been cancelled or not

```go
ctx := context.Background()
resultsCh := make(chan *WorkResult)

for {
	select {
	case <- ctx.Done():
		// The context is over, stop processing results
		return
	case result := <- resultsCh:
		// Process the results received
	}
}
```

### Cancellign a context

canccling a context is done by calling the cancel function returned by the context.WithCancel function

```go
ctx, cancel := context.WithCancel(context.Background())
cancel()
```

If you’ve run Go programs before and looked at the logging output, you may have seen the context canceled error in the past. When using the Go http package, this is a common error to see when a client disconnects from the server before the server handles the full response

eg example.go


## Context with deadline 

- Using context.withdeadline with a context allows you to set a deadline for when the context needs to be finishhed and it automatically ends when deadline passed

- When a context is canceled from a deadline, the cancel function is still required to be called in order to clean up any resources that were used, so this is more of a safety measure


## Context with timeout

- Same as context with deadline but instead of deadline it takes a timeout