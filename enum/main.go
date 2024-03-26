package main

import "fmt"

type WeaponType int

const (
	Axe WeaponType = iota
	Sword
	WoodenSword
	Knife
)

func (wt WeaponType) String() string {
	switch wt {
	case Axe:
		return "AXE"
	case Sword:
		return "SWORD"
	case WoodenSword:
		return "WOODEN SWORD"
	case Knife:
		return "KNIFE"
	default:
		panic("Unknown weapon type")
	}
}

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 10
	case Sword:
		return 15
	case WoodenSword:
		return 5
	case Knife:
		return 3
	default:
		panic("Unknown weapon type")
	}
}

func main() {
	for _, weaponType := range []WeaponType{Axe, Sword, WoodenSword, Knife} {
		fmt.Printf("%s: %d\n", weaponType, getDamage(weaponType))
	}

}
