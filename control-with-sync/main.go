package main

import (
	"fmt"
	"time"
)

func main() {
	server := newServer()

	go func() {
		for i := 0; i < 100; i++ {
			server.send(fmt.Sprintf("Message %d", i))
		}

		time.Sleep(1 * time.Second)
		server.quit()
	}()

	server.start()
}

type Server struct {
	quitchan chan struct{}
	msgchan  chan string
}

func newServer() *Server {
	return &Server{
		quitchan: make(chan struct{}),
		msgchan:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("Start")
	s.loop()
}

func (s *Server) send(msg string) {
	s.msgchan <- msg
}

func (s *Server) quit() {
	s.quitchan <- struct{}{}
}

func (s *Server) loop() {
out: // Label for the for loop
	for {
		select {
		case <-s.quitchan:
			fmt.Println("Quit")
			break out
		case msg := <-s.msgchan:
			fmt.Println(msg)
		default:
		}
	}
}
