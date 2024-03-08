package main

import (
	"fmt"
	"time"
)

// Buffered channels are used to store data in a queue. The data is stored in a first-in-first-out order.

func main() {
	fmt.Println("Start")
	done := make(chan bool) // Unbuffered channel (Always use unbuffered channels when you need to synchronize the execution of goroutines)

	fmt.Println("non-concurrent call")
	go externalCall("Non-Concurrent", done)

	fmt.Println("concurrent call")
	externalCall("Concurrent", done)
	done <- true // Blocks until the value is read from the channel

	fmt.Println("Done")
}

func externalCall(callType string, blockChan <-chan bool) string {
	now := time.Now()
	if callType == "Concurrent" {
		time.Sleep(1 * time.Second)
	} else {
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("externalCall %v took %v\n", callType, time.Since(now))
	if callType != "Concurrent" {
		<-blockChan
	}
	return "Hello, World"
}
