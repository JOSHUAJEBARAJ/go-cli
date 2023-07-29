package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// this is sequential program took 12 sec
	fmt.Println(runtime.GOMAXPROCS(0))
	//runtime.GOMAXPROCS(8)
	start := time.Now()
	counta()
	countb()
	countc()
	countd()
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}

func counta() {
	fmt.Println("AAA Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("A is done")
}

func countb() {
	fmt.Println("BBB Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("B is done")
}

func countc() {
	fmt.Println("CCC Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("C is done")
}

func countd() {
	fmt.Println("DDD Starting")
	for i := 1; i < 10_000_000_000; i++ {

	}
	fmt.Println("D is done")
}
