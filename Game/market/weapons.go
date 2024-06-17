package market

import (
	"fmt"
)

var Itemdamage int

func Weapon_List() {
	var loop bool
	var decision int
	fmt.Printf("Welcome to weoponsmith warrior\n Tier1: Wooden Sword; 25 damage; 50 gold \n Tier2: Copper sword; 50 damage; 100 gold \n Tier3 Iron sword; 75 damage; 250 gold \n Tier4: Obsidian Sword; 120 damage; 500 gold \n You have %v gold \n 5. Main menu\n", Gold)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			if Gold >= 50 {
				Gold = Gold - 50
				Itemdamage = 25
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You dont have enough money to buy wooden armour")
			}
		case decision == 2:
			if Gold >= 100 {
				Gold = Gold - 100
				Itemdamage = 50
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You don't have enough money to buy copper armour")
			}
		case decision == 3:
			if Gold >= 250 {
				Gold = Gold - 250
				Itemdamage = 75
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You don't have enough money to buy iron armour")
			}
		case decision == 4:
			if Gold >= 500 {
				Gold = Gold - 500
				Itemdamage = 90
				loop = true
				fmt.Println("Thank you for your purchase")
				break
			} else {
				fmt.Println("\n You don't have enough money to buy chainmail armour")
			}
		case decision == 5:
			loop = true
			break
		default:
			fmt.Println("\nPlease choose one of the options\n")
		}
	}
}
