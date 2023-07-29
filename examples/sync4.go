package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			var result = 0
			goChan := make(chan int)
			mainChan := make(chan string)
			calculateSquare := func() {
				fmt.Println("performing some operation") // these wait for three seconds
				time.Sleep(time.Second * 3)
				result = j * j
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
		}(i)
		wg.Wait()
	}
}
