package main

import (
	"context"
	"fmt"
)

func main() {
	//ctx := context.TODO()
	ctx := context.Background()
	doSomething(ctx)

}

func doSomething(ctx context.Context) {
	fmt.Println("Hello world")
}
