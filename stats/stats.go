package stats

import (
	"fmt"

	"github.com/M1ralai/RPG/market"
)

var Hp, Damage, Armour int

func Stats() {
	Hp = (150) + (50 * Level)
	Damage = (30 + market.Itemdamage) + (30 * Level) + (Level*market.Itemdamage)/2
	Armour = (15 + market.Itemarmour) + (5 * Level) + (Level*market.Itemarmour)/2
}
func Stats_Visualisation() {
	var decision int
	var loop bool
	fmt.Printf("\n   Your hp: %v \n   Your armour: %v \n   Your damage: %v \n", Hp, Armour, Damage)
	fmt.Println(" \n 1. Return main menu \n\n 2.Quit")
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		if decision == 1 {
			loop = true
		} else if decision == 2 {
			loop = true
			return
		} else {
			fmt.Println("Please choose one of the options.")
		}
	}
}
