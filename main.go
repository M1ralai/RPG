package main

import (
	"fmt"
)

var level, decision, gold int
var loop bool

func main() {
	fmt.Println("Welcome to the real world warrior choose where you wanna go \n\n 1.Marketplace \n\n 2.Monsters Hub")
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		if decision == 1 {
			loop = true
			marketplace()
		} else if decision == 2 {
			loop = true
			monster_hub()
		} else {
			fmt.Println("Please choose one of the options")
		}
	}
}
func marketplace() {
	fmt.Println("Do you comfirm to go Marketplace?\n 1.Yes\n 2.No get me back to Main Menu")
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		if decision == 1 {
			loop = true
			market_list()
		} else if decision == 2 {
			main()
			loop = true
		} else {
			fmt.Println("Please select one of the options")
		}
	}
}

func monster_hub() {
	fmt.Println("Do you confirm to go Monsters Hub? \n 1.Yes\n  2.No get me back to Main Menu")
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		if decision == 1 {
			loop = true
			monsters()
		} else if decision == 2 {
			main()
			loop = true
		} else {
			fmt.Println("Please choose one of the options")
		}

	}
}
func monsters() {
	fmt.Println("You entered this endless space, so tell me qich one do you wanna fight?\n 1. Scarab\n 2. Zombie\n 3. Skeleton ")
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		switch {
		case decision == 1:
			fmt.Println("You chhosed Scarab...")
			loop = true
			fight_scarab()
		case decision == 2:
			fmt.Println("You choosed Zombie...")
			fight_zombie()
			loop = true
		case decision == 3:
			fmt.Println("You choosed Skeleton...")
			fight_skeleton()
			loop = true
		default:
			fmt.Println("Please choose one of the options")
		}
	}
}
func fight_scarab() {
	var scarab_hp, scarab_armour, scarab_damage int
	scarab_hp = 100 * ((150 / 100) ^ level)
	scarab_armour = 50 * ((125 / 100) ^ level)
	scarab_damage = 25 * ((175 / 100) ^ level)
	fmt.Printf("\n  Scarabs hp: %v \n  Scarabs armour: %v \n  Scarabs Damage: %v \n\n Do you accept this fight? \n\n 1.Yes \n\n 2.No return monster hub \n\n 3.No return main menu\n", scarab_hp, scarab_armour, scarab_damage)
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		switch {
		case decision == 1:

		case decision == 2:
			monster_hub()
			loop = true
		case decision == 3:
			main()
			loop = true
		default:
			fmt.Println("Please choose one od the options.")
		}
	}
}
func fight_zombie() {
	var zombie_hp, zombie_armour, zombie_damage int
	zombie_hp = 150 * ((115 / 100) ^ level)
	zombie_armour = 100 * ((150 / 100) ^ level)
	zombie_damage = 100 * ((115 / 100) ^ level)
	fmt.Printf("\n Zombies hp: %v \n Zombies armour: %v \n Zombies damage: %v \n\n Do you accept this fihgt? \n\n 1.Yes \n\n 2.No return monster hub \n\n 3.No return main menu\n", zombie_hp, zombie_armour, zombie_damage)
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		switch {
		case decision == 1:

		case decision == 2:
			monster_hub()
			loop = true
		case decision == 3:
			main()
			loop = true
		default:
			fmt.Println("Please choose one od the options.")
		}
	}

}
func fight_skeleton() {
	var skeleton_hp, skeleton_armour, skeleton_damage int
	skeleton_hp = 50 * ((125 / 100) ^ level)
	skeleton_armour = 100 * ((150 / 100) ^ level)
	skeleton_damage = 50 * ((200 / 100) ^ level)
	fmt.Printf("\n Skeletons hp: %v \n Skeletons armour: %v \n Skeletons damage: %v \n\n Do you accept this fight? \n\n 1.Yes \n\n 2.No return monster hub \n\n 3.No return main menu\n", skeleton_hp, skeleton_armour, skeleton_damage)
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		switch {
		case decision == 1:

		case decision == 2:
			monster_hub()
			loop = true
		case decision == 3:
			main()
			loop = true
		default:
			fmt.Println("Please choose one od the options.")
		}
	}
}
func market_list() {
	fmt.Println("Welcome to the market sir, here is some upgrades for you \n\n  1.Armour list \n 2.Weapon list")
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		switch {
		case decision == 1:
			armour_list()
			loop = true
		case decision == 2:
			weapon_list()
			loop = true
		default:
			fmt.Println("Please chose one of the options.")
		}
	}
}
func armour_list() {
	fmt.Println("Welcome to armoursmith warrior.")
}
func weapon_list() {
	fmt.Println("Welcome to weoponsmith warrior")
}
