package main

import "fmt"

func main() {
	c := make(chan string, 3)
	c <- "one"
	c <- "two"
	c <- "three"
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
