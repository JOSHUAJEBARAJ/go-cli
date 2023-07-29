package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndTalk(ctx, 5*time.Second, "Hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {

	select {
	case <-time.After(d):
		fmt.Print(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
