package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var level, decision, gold, hp, damage, armour, itemarmour, itemdamage, rng, hit1, hit2, scarab_counter, exp, zombie_counter, skeleton_counter int
var loop bool

func main() {
	fmt.Println("Welcome to the real world warrior choose where you wanna go \n\n 1.Marketplace \n\n 2.Monsters Hub \n\n 3. Show stats\n")
	loop = false
	hp = (150) + (50 * level)
	damage = (30 + itemdamage) + (30 * level) + (level*itemdamage)/2
	armour = (15 + itemarmour) + (5 * level) + (level*itemarmour)/2
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			market_list()
			loop = true
		case decision == 2:
			monsters()
			loop = true
		case decision == 3:
			stats()
			loop = true
		default:
			fmt.Println("Please choose one of the options.")
		}
	}
}
func stats() {
	hp = (150) + (50 * level)
	damage = (30 + itemdamage) + (30 * level) + (level*itemdamage)/2
	armour = (15 + itemarmour) + (5 * level) + (level*itemarmour)/2
	fmt.Printf("\n   Your hp: %v \n   Your armour: %v \n   Your damage: %v \n", hp, armour, damage)
	fmt.Println(" \n 1. Return main menu \n\n 2.Quit")
	fmt.Scanln(&decision)
	loop = false
	for loop == false {
		if decision == 1 {
			loop = true
			main()
		} else if decision == 2 {
			loop = true
			return
		} else {
			fmt.Println("Please choose one of the options.")
		}
	}
}
func monsters() {
	fmt.Println("You entered this endless space, so tell me qich one do you wanna fight?\n 1. Scarab\n 2. Zombie\n 3. Skeleton \n 4.Return menu\n")
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
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
		case decision == 4:
			main()
			loop = true
		default:
			fmt.Println("Please choose one of the options")
		}
	}
}
func fight_scarab() {
	var scarab_hp, scarab_armour, scarab_damage, scarab_gold, scarab_exp int
	scarab_hp = 100 + (20 * level)
	if scarab_armour <= 100 {
		scarab_armour = 50 + (10 * level)
	}
	scarab_damage = 25 + (15 * level)
	scarab_gold = 20 + (10 * level)
	scarab_exp = 50 + (10 * level)

	fmt.Printf("\n  Scarabs hp: %v \n  Scarabs armour: %v \n  Scarabs Damage: %v \n\n Do you accept this fight? \n\n 1.Yes \n\n 2.No return monster hub \n\n 3.No return main menu\n", scarab_hp, scarab_armour, scarab_damage)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			for hp >= 0 && scarab_hp >= 0 {
				if armour >= scarab_damage {
					fmt.Println("You are too strong for this monster; You have to go main menu...")
					main()
				}
				rng = rand.IntN(40)
				duration := time.Second
				time.Sleep(duration)
				hit1 = (damage * (80 + rng) / 100) * (150 - scarab_armour) / 150
				scarab_hp = scarab_hp - hit1
				rng = rand.IntN(40)
				hit2 = (scarab_damage * (80 + rng) / 100) * (150 - armour) / 150
				hp = hp - hit2
				if hp >= 0 && scarab_hp >= 0 {
					fmt.Printf("\n\n Scarab attacked you and hit: %v \n You have %v hp now. \n\n", hit2, hp)
					fmt.Printf(" \n\n Your attacked and hit: %v \n Scarab have %v hp now. \n\n", hit1, scarab_hp)
				} else if hp >= 0 {
					fmt.Println(" \n \nYou win!!! \n \n")
					exp = exp + scarab_exp
					scarab_counter += 1
					fmt.Printf(" You earned %v experience;\n\n You gained %v gold ", exp, scarab_gold)
					gold = gold + scarab_gold
					exp_level()
				} else {
					fmt.Println("\n\nYou died but people gave you 10 gold for honor.\n\n")
					gold = gold + 10
					main()
				}
			}
		case decision == 2:
			monsters()
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
	var zombie_hp, zombie_armour, zombie_damage, zombie_gold, zombie_exp int
	zombie_hp = 150 + (30 * level)
	zombie_armour = 100
	zombie_damage = 100 + (30 * level)
	zombie_exp = 75 + (50 * level)
	zombie_gold = 50 + (50 * level)
	fmt.Printf("\n Zombies hp: %v \n Zombies armour: %v \n Zombies damage: %v \n\n Do you accept this fihgt? \n\n 1.Yes \n\n 2.No return monster hub \n\n 3.No return main menu\n", zombie_hp, zombie_armour, zombie_damage)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			for hp >= 0 && zombie_hp >= 0 {
				if armour >= zombie_damage {
					fmt.Println("You are too strong for this monster; You have to go main menu...")
					main()
				}
				rng = rand.IntN(40)
				duration := time.Second
				time.Sleep(duration)
				hit1 = (damage * (80 + rng) / 100) * (150 - zombie_armour) / 150
				zombie_hp = zombie_hp - hit1
				rng = rand.IntN(40)
				hit2 = (zombie_damage * (80 + rng) / 100) * (150 - armour) / 150
				hp = hp - hit2
				if hp >= 0 && zombie_hp >= 0 {
					fmt.Printf("\n\n Scarab attacked you and hit: %v \n You have %v hp now. \n\n", hit2, hp)
					fmt.Printf(" \n\n Your attacked and hit: %v \n Scarab have %v hp now. \n\n", hit1, zombie_hp)
				} else if hp >= 0 {
					fmt.Println(" \n \nYou win!!! \n \n")
					exp = exp + zombie_exp
					scarab_counter += 1
					fmt.Printf(" You earned %v experience;\n\n You gained %v gold ", exp, zombie_gold)
					gold = gold + zombie_gold
					exp_level()
				} else {
					fmt.Println("\n\nYou died but people gave you 25 gold for honor.\n\n")
					gold = gold + 25
					main()
				}
			}
		case decision == 2:
			monsters()
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
	var skeleton_hp, skeleton_armour, skeleton_damage, skeleton_exp, skeleton_gold int
	skeleton_hp = 50 * ((125 / 100) ^ level)
	skeleton_armour = 100 * ((150 / 100) ^ level)
	skeleton_damage = 50 * ((200 / 100) ^ level)
	skeleton_exp = 100 + (75 * level)
	skeleton_gold = 100 + (75 * gold)
	fmt.Printf("\n Skeletons hp: %v \n Skeletons armour: %v \n Skeletons damage: %v \n\n Do you accept this fight? \n\n 1.Yes \n\n 2.No return monster hub \n\n 3.No return main menu\n", skeleton_hp, skeleton_armour, skeleton_damage)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			for hp >= 0 && skeleton_hp >= 0 {
				if armour >= skeleton_damage {
					fmt.Println("You are too strong for this monster; You have to go main menu...")
					main()
				}
				rng = rand.IntN(40)
				duration := time.Second
				time.Sleep(duration)
				hit1 = (damage * (80 + rng) / 100) * (150 - skeleton_armour) / 150
				skeleton_hp = skeleton_hp - hit1
				rng = rand.IntN(40)
				hit2 = (skeleton_damage * (80 + rng) / 100) * (150 - armour) / 150
				hp = hp - hit2
				if hp >= 0 && skeleton_hp >= 0 {
					fmt.Printf("\n\n Scarab attacked you and hit: %v \n You have %v hp now. \n\n", hit2, hp)
					fmt.Printf(" \n\n Your attacked and hit: %v \n Scarab have %v hp now. \n\n", hit1, skeleton_hp)
				} else if hp >= 0 {
					fmt.Println(" \n \nYou win!!! \n \n")
					exp = exp + skeleton_exp
					scarab_counter += 1
					fmt.Printf(" You earned %v experience;\n\n You gained %v gold ", exp, skeleton_gold)
					gold = gold + skeleton_gold
					exp_level()
				} else {
					fmt.Println("\n\nYou died but people gave you 25 gold for honor.\n\n")
					gold = gold + 25
					main()
				}
			}
		case decision == 2:
			monsters()
			loop = true
		case decision == 3:
			main()
			loop = true
		default:
			fmt.Println("Please choose one of the options.")
		}
	}
}
func market_list() {
	fmt.Printf("Welcome to the market sir, here is some items for you \n\n 1.Armour list \n 2.Weapon list \n  3.Return main menu \n You have %v gold\n", gold)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			armour_list()
			loop = true
		case decision == 2:
			weapon_list()
			loop = true
		case decision == 3:
			main()
			loop = true
		default:
			fmt.Println("Please chose one of the options.")
		}
	}
}
func armour_list() {
	fmt.Printf("Welcome to armoursmith warrior. Here is some armours. \n Tier1: Wooden armour; 25 armour; 50 gold \n Tier2: Copper armour; 50 armour; 100 gold \n Tier3: Iron armour; 75 armour; 250 gold \n Tier4: Chainmail Armour; 90 armour; 500 Gold \n You have %v gold \n 5. Main menu\n", gold)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			if gold >= 50 {
				gold = gold - 50
				itemarmour = 25
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You dont have enough money to buy wooden armour")
			}
		case decision == 2:
			if gold >= 100 {
				gold = gold - 100
				itemarmour = 50
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You don't have enough money to buy copper armour")
			}
		case decision == 3:
			if gold >= 250 {
				gold = gold - 250
				itemarmour = 75
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You don't have enough money to buy iron armour")
			}
		case decision == 4:
			if gold >= 500 {
				gold = gold - 500
				itemarmour = 90
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You don't have enough money to buy chainmail armour")
			}
		case decision == 5:
			main()
		default:
			fmt.Println("\nPlease choose one of the options\n")
		}
	}
}
func weapon_list() {
	fmt.Printf("Welcome to weoponsmith warrior\n Tier1: Wooden Sword; 25 damage; 50 gold \n Tier2: Copper sword; 50 damage; 100 gold \n Tier3 Iron sword; 75 damage; 250 gold \n Tier4: Obsidian Sword; 120 damage; 500 gold \n You have %v gold \n 5. Main menu\n", gold)
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			if gold >= 50 {
				gold = gold - 50
				itemdamage = 25
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You dont have enough money to buy wooden armour")
			}
		case decision == 2:
			if gold >= 100 {
				gold = gold - 100
				itemdamage = 50
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You don't have enough money to buy copper armour")
			}
		case decision == 3:
			if gold >= 250 {
				gold = gold - 250
				itemdamage = 75
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You don't have enough money to buy iron armour")
			}
		case decision == 4:
			if gold >= 500 {
				gold = gold - 500
				itemdamage = 90
				loop = true
				fmt.Println("Thank you for your purchase")
				stats()
			} else {
				fmt.Println("\n You don't have enough money to buy chainmail armour")
			}
		case decision == 5:
			main()
		default:
			fmt.Println("\nPlease choose one of the options\n")
		}
	}
}
func exp_level() {
	var remain, cap int
	cap = 100 + (level * 200)
	remain = cap - exp
	if exp < cap {
		fmt.Printf("\n\nYou have %v experiende and you need to earn %v experience top level up\n\n", exp, remain)
		fmt.Println("1.Main menu")
	} else {
		level += 1
		fmt.Printf("\n\nCongrats you leveled up and you %v level now\n\n", level)
		exp = exp - cap
		scarab_counter = 0
		zombie_counter = 0
		skeleton_counter = 0
	}
	loop = false
	for loop == false {
		fmt.Scanln(&decision)
		if decision == 1 {
			loop = true
			main()
		} else {
			fmt.Println("Please choose one of the options.")
		}
	}
}
