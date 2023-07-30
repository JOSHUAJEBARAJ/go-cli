package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	printCh := make(chan int)
	go func() {
		doAnother(ctx, printCh)
	}()
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	cancel()
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
