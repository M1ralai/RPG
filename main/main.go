package main

import (
	"fmt"
)

func main() {
	var loop bool
	var decision int
	fmt.Println("Welcome to the real world warrior choose where you wanna go \n\n 1.Marketplace \n\n 2.Monsters Hub")
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		if decision == 1 {
			loop = true
			marketplace(decision, loop)
		} else if decision == 2 {
			loop = true
			monster_hub(decision, loop)
		} else {
			fmt.Println("Please choose one of the options")
		}
	}
}
func marketplace(decision int, loop bool) {
	fmt.Println("Do you comfirm to go Marketplace? \n 1.Yes \n 2.No get me back to Main Menu")
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		if decision == 1 {
			//Send man/women to a market.go
			loop = true
		} else if decision == 2 {
			main()
			loop = true
		} else {
			fmt.Println("Please select one of the options")
		}
	}
}

func monster_hub(decision int, loop bool) {
	fmt.Println("Do you confirm to go Monsters Hub? \n 1.Yes\n  2.No get me back to Main Menu")
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		if decision == 1 {
			//Send man/women to a monsters.go
			loop = true
		} else if decision == 2 {
			main()
			loop = true
		} else {
			fmt.Println("Please choose one of the options")
		}

	}
}
