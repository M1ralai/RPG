package market

import (
	"fmt"
)

var Itemarmour int

func armour_list() {
	var loop bool
	var decision int
	fmt.Printf("Welcome to armoursmith warrior. Here is some armours. \n Tier1: Wooden armour; 25 armour; 50 gold \n Tier2: Copper armour; 50 armour; 100 gold \n Tier3: Iron armour; 75 armour; 250 gold \n Tier4: Chainmail Armour; 90 armour; 500 Gold \n You have %v gold \n 5. Main menu\n", Gold)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			if Gold >= 50 {
				Gold = Gold - 50
				Itemarmour = 25
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You dont have enough money to buy wooden armour")
			}
		case decision == 2:
			if Gold >= 100 {
				Gold = Gold - 100
				Itemarmour = 50
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You don't have enough money to buy copper armour")
			}
		case decision == 3:
			if Gold >= 250 {
				Gold = Gold - 250
				Itemarmour = 75
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You don't have enough money to buy iron armour")
			}
		case decision == 4:
			if Gold >= 500 {
				Gold = Gold - 500
				Itemarmour = 90
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You don't have enough money to buy chainmail armour")
			}
		case decision == 5:
			break
		default:
			fmt.Println("\nPlease choose one of the options\n")
		}
	}
}
