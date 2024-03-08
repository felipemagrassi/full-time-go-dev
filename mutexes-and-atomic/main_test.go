package main

import (
	"sync"
	"testing"
)

func TestStateAtomic(t *testing.T) {
	state := &State{}
	wg := &sync.WaitGroup{}
	sum := 0

	for i := 0; i < 100; i++ {
		sum += i

		wg.Add(1)
		go func(i int) {
			state.atomicSet(i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	if state.count != int32(sum) {
		t.Errorf("State count is %d, want %d", state.count, int32(sum))
	}

}

func TestStateMutex(t *testing.T) {
	state := &State{}
	wg := &sync.WaitGroup{}
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
		wg.Add(1)
		go func(i int) {
			state.set(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	if state.count >= 100 {
		t.Errorf("State count is %d, want 100", state.count)
	}
}
