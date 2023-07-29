package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// this is sequential program took 12 sec
	fmt.Println(runtime.GOMAXPROCS(0))
	//runtime.GOMAXPROCS(8)
	start := time.Now()
	c := make(chan string)
	go counta(c)
	go countb(c)
	go countc(c)
	go countd(c)
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}

func counta(c chan string) {
	fmt.Println("AAA Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("A is done")
	c <- "A Done"
}

func countb(c chan string) {
	fmt.Println("BBB Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("B is done")
	c <- "B Done"
}

func countc(c chan string) {
	fmt.Println("CCC Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("C is done")
	c <- "C Done"
}

func countd(c chan string) {
	fmt.Println("DDD Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("D is done")
	c <- "D Done"
}
