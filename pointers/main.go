package main

import "fmt"

type Player struct {
	HP int
}

func TakeDamage(player Player, amount int) {
	fmt.Printf("Address of player = %+v, %p\n", player, &player)
	fmt.Printf("Player hp before damage %d\n", player.HP)
	player.HP -= amount
	fmt.Printf("Player took %d damage, HP went down to %d\n", amount, player.HP)
}

func TakeDamageByPointer(player *Player, amount int) {
	fmt.Printf("Address of player = %+v, %p\n", player, player)
	fmt.Printf("Address of player = %+v, %p\n", player, &player)
	fmt.Printf("Player hp before damage %d\n", player.HP)
	player.HP -= amount
	fmt.Printf("Player took %d damage, HP went down to %d\n", amount, player.HP)
}

func main() {
	player := Player{
		HP: 100,
	}
	fmt.Printf("Address of player = %+v, %p\n", player, &player)
	fmt.Printf("----------------------\n")

	TakeDamage(player, 10)
	fmt.Printf("Player hp after function call %d\n", player.HP)
	fmt.Printf("----------------------\n")

	TakeDamageByPointer(&player, 10)

	fmt.Printf("Player hp after pointer function call %d\n", player.HP)
	fmt.Printf("----------------------\n")
}
