package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	start := time.Now()
	go sayHello()
	go sayBye()
	wg.Wait()

	fmt.Printf("proccess took &%s", time.Since(start))

}

func sayHello() {
	time.Sleep(time.Second * 2)
	fmt.Println("Hello")
	wg.Done()
}

func sayBye() {
	time.Sleep(time.Second * 2)
	fmt.Println("Bye")
	wg.Done()
}
