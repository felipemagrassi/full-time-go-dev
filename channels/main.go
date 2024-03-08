package main

import (
	"fmt"
	"time"
)

func main() {
	// firstExample() Consumer doesnt know when to stop reading from the channel
	// secondExample() // Channels closes and producer knows when to stop reading
	thirdExample() // same as secondExample but with a break statement
}

func thirdExample() {
	msgCh := make(chan string, 128)

	msgCh <- "Hello"
	msgCh <- "World"
	msgCh <- "!"
	close(msgCh)

	for {
		msg, ok := <-msgCh
		if !ok {
			break
		}

		fmt.Println("Message:", msg)
	}
}

func secondExample() {
	msgCh := make(chan string, 128)

	msgCh <- "Hello"
	msgCh <- "World"
	msgCh <- "!"
	close(msgCh)

	for msh := range msgCh {
		fmt.Printf(msh)
	}
}

func firstExample() {
	msgCh := make(chan string, 128)

	msgCh <- "Hello"
	msgCh <- "World"
	msgCh <- "!"

	for msh := range msgCh {
		fmt.Println(msh)
	}
}

func externalCall() {
	time.Sleep(2 * time.Second)
}
