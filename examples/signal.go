package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time" // or "runtime"
)

func cleanup() {
	fmt.Println("cleanup")
}

func main() {
	c := make(chan os.Signal)
	// The signal.Notify function is called to register the channel c to receive specific signals: os.Interrupt (Ctrl+C) and syscall.SIGTERM (terminate signal). When any of these signals are received, they will be sent on the channel c.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}
