package main

import "fmt"

type Position struct {
	x, y int
}

func (p *Position) Move(dx, dy int) {
	p.x += dx
	p.y += dy
}

type Snake struct {
	Position
}

type Color int

func (c Color) String() string {
	return [...]string{"#0000FF", "#00FF00", "#000000", "#FFFF00"}[c]
}

const (
	Blue Color = iota
	Green
	Black
	Yellow
)

func main() {
	snake := Snake{
		Position{x: 10, y: 10},
	}
	fmt.Println(snake.x, snake.y)

	snake.x = 20
	snake.y = 20

	fmt.Println(snake.x, snake.y)
	fmt.Println(snake.Position.x, snake.Position.y)

	snake.Move(5, 5)
	fmt.Println(snake.x, snake.y)
	fmt.Println(snake.Position.x, snake.Position.y)

	fmt.Println(Blue)
	fmt.Println(Black)

}
