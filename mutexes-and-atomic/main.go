package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type State struct {
	count int32
	mu    sync.RWMutex
}

func (s *State) set(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = int32(i)
}

func (s *State) atomicSet(i int) {
	atomic.AddInt32(&s.count, int32(i))
}

func main() {
	state := &State{}

	for i := 0; i < 1000; i++ {
		state.count = int32(i)
	}

	fmt.Println(state.count)
}
