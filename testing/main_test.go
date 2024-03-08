package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculateValue(t *testing.T) {
	fmt.Println("Verbose Mode")
	// t.Fail() -- force fail

	var (
		expected = 5
		a        = 2
		b        = 3
	)

	have := CalculateValue(a, b)
	if have != expected {
		t.Errorf("Expected %v but got %v", expected, have)
	}

}

func TestCalculateValueSpecial(t *testing.T) {
	player1 := Player{
		HP:   100,
		Name: "Player 1",
	}

	player2 := Player{
		HP:   100,
		Name: "Player 1",
	}

	if !reflect.DeepEqual(player1, player2) {
		t.Errorf("Expected %+v but got %+v", player1, player2)
	}
}
