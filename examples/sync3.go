package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var result = 0
		var value = 3
		goChan := make(chan int)
		mainChan := make(chan string)
		calculateSquare := func() {
			fmt.Println("performing some operation") // these wait for three seconds
			time.Sleep(time.Second * 3)
			result = value * value
			goChan <- result
		}
		reportResult := func() {
			time.Sleep(time.Second * 1)
			fmt.Println("The result is ", <-goChan)
			mainChan <- "Done"

		}
		go calculateSquare()
		go reportResult()
		<-mainChan
		wg.Done()
	}()
	wg.Wait()
}
