package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {
	start := time.Now()
	go sayHello()
	go sayBye()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Printf("proccess took &%s", time.Since(start))

}

func sayHello() {
	time.Sleep(time.Second * 2)
	// fmt.Println("Hello")
	ch <- "SayHello"
}

func sayBye() {
	time.Sleep(time.Second * 2)
	// fmt.Println("Bye")
	ch <- "Bye"
}
