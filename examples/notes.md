## Concurrency 

Concurrency is the combination of the parallelism and multi tasking 

## Pattern 
Instead of running the function prefix it with go keyword

```go
go test()
```

It's always to wrap it with the anonymous function

```go
go func(){
test()
}()

## Go routines

- Light weight thread managed by go runtime

## Wait groups

```go
var wg = sync.Waitgroup{}
wg.Add(1)
wg.Done()
wg.Wait()
```

## Channels 

- Channels are used to communicate between go routines and synchronize them
- Don't communicate by sharing memory, share memory by communicating

```go
ch := make(chan string)
ch <- "Hello"(BLOCKING)
msg := <- ch(BLOCKING)
close(ch)
```

## Buffered channels

```go
ch := make(chan string, 3)
```

Works in FIFO order

## Using go to synchronize

## Io-Bound vs CPU-Bound

- I/O bound processes have performance limited by input/output operations
- CPU bound processes have performance limited by CPU speed

