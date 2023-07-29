package main

import (
	"fmt"
	"time"
)

var result = 0
var value = 3

func main() {
	goChan := make(chan int)
	mainChan := make(chan string)
	go calculateSquare(value, goChan)
	go reportResult(goChan, mainChan)
	<-mainChan
}

func calculateSquare(value int, goChan chan int) {
	fmt.Println("performing some operation") // these wait for three seconds
	time.Sleep(time.Second * 3)
	result = value * value
	goChan <- result
}

func reportResult(goChan chan int, mainChan chan string) {
	time.Sleep(time.Second * 1)
	fmt.Println("The result is ", <-goChan)
	mainChan <- "Done"
}
