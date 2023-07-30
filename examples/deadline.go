package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	printCh := make(chan int)
	go func() {
		doAnother(ctx, printCh)
	}()
	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}
	cancel() //When a context is canceled from a deadline, the cancel function is still required to be called in order to clean up any resources that were used, so this is more of a safety measure
	time.Sleep(2 * time.Second)
	fmt.Printf("Do something done")
}

func doAnother(ctx context.Context, printch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err %s\n", err)
			}
			fmt.Printf("Do another done\n")
			return
		case num := <-printch:
			fmt.Println(num)
		}
	}
}

func main() {
	//ctx := context.TODO()
	ctx := context.Background()
	doSomething(ctx)

}
