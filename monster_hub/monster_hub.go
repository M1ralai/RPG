package monster_hub

import (
	"fmt"
)

var Monster_Decision int

func Monsters() {
	var loop bool
	var decision int
	fmt.Println("You entered this endless space, so tell me qich one do you wanna fight?\n 1. Scarab\n 2. Zombie\n 3. Skeleton \n 4.Return menu\n")
	for !loop {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			fmt.Println("You chhosed Scarab...")
			loop = true
			Monster_Decision = 1
			Fight()
		case decision == 2:
			fmt.Println("You choosed Zombie...")
			Monster_Decision = 2
			Fight()
			loop = true
		case decision == 3:
			fmt.Println("You choosed Skeleton...")
			Monster_Decision = 3
			Fight()
			loop = true
		case decision == 4:
			loop = true
		default:
			fmt.Println("Please choose one of the options")
		}
	}
}
