package main

import (
	"context"
	"fmt"
)

func main() {
	//ctx := context.TODO()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "joshua")
	doSomething(ctx)

}

func doSomething(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
}
