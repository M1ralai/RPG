package market

import "fmt"

var Gold int

func Market_Section() {
	var loop bool
	var decision int
	fmt.Printf("Welcome to the market sir, here is some items for you \n\n 1.Armour list \n 2.Weapon list \n 3.Return main menu \n You have %v gold\n", Gold)
	loop = false
	for !loop {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			armour_list()
			loop = true
		case decision == 2:
			Weapon_List()
			loop = true
		case decision == 3:
			loop = true
		default:
			fmt.Println("Please chose one of the options.")
		}
	}
}
