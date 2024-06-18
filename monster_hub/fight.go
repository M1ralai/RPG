package monster_hub

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/M1ralai/RPG/market"
	"github.com/M1ralai/RPG/stats"
)

func Fight() {
	var decision, monster_damage, monster_armour, monster_hp, monster_exp, monster_gold, rng, hit1, hit2 int
	var loop bool
	var monster_name string
	scarab := monsters{"Scarab", 100, 25, 20, 50, 50}
	zombie := monsters{"Zombie", 150, 100, 50, 75, 100}
	skeleton := monsters{"Skeleton", 150, 100, 50, 75, 100}
	if Monster_Decision == 1 {
		monster_hp = scarab.hp + (20 * stats.Level)
		monster_armour = scarab.armour + (10 * stats.Level)
		if monster_armour >= 100 {
			monster_armour = 100
		}
		monster_gold = scarab.gold + (stats.Level * 20)
		monster_damage = scarab.damage + (15 * stats.Level)
		monster_exp = scarab.exp + (10 * stats.Level)
		monster_name = scarab.name
	} else if Monster_Decision == 2 {
		monster_hp = zombie.hp + (30 * stats.Level)
		monster_damage = zombie.damage + (30 * stats.Level)
		monster_gold = zombie.gold + (50 * stats.Level)
		monster_exp = zombie.exp + (50 * stats.Level)
		monster_armour = zombie.armour
		monster_name = "Zombie"
	} else if Monster_Decision == 3 {
		monster_hp = skeleton.hp + (30 * stats.Level)
		monster_damage = 100 + (30 * stats.Level)
		monster_gold = 50 + (50 * stats.Level)
		monster_exp = 75 + (50 * stats.Level)
		monster_armour = 100
		monster_name = "Skeleton"
	}
	fmt.Printf("\n  %v hp: %v \n  %v armour: %v \n  %v Damage: %v \n\n Do you accept this fight? \n\n 1.Yes \n\n 2.No return monster hub \n\n 3.No return main menu\n", monster_name, monster_name, monster_name, monster_hp, monster_armour, monster_damage)
	loop = false
	if stats.Armour >= monster_damage {
		fmt.Println("You are too strong for this monster; You have to go main menu...")

	} else {
		for !loop {
			fmt.Scanln(&decision)
			switch {
			case decision == 1:
				for stats.Hp >= 0 && monster_hp >= 0 {
					rng = rand.IntN(40)
					duration := time.Second
					time.Sleep(duration)
					hit1 = (stats.Damage * (80 + rng) / 100) * (150 - monster_armour) / 150
					monster_hp = monster_hp - hit1
					rng = rand.IntN(40)
					hit2 = (monster_damage * (80 + rng) / 100) * (150 - stats.Armour) / 150
					stats.Hp = stats.Hp - hit2
					if stats.Hp >= 0 && monster_hp >= 0 {
						fmt.Printf("\n\n %v attacked you and hit: %v \n You have %v hp now. \n\n", monster_name, hit2, stats.Hp)
						fmt.Printf(" \n\n Your attacked and hit: %v \n %v have %v hp now. \n\n", hit1, monster_name, monster_hp)
					} else if stats.Hp >= 0 {
						fmt.Println(" \n \nYou win!!! \n \n")
						stats.Exp = stats.Exp + monster_exp
						fmt.Printf(" You earned %v experience;\n\n You gained %v gold ", stats.Exp, monster_gold)
						market.Gold = market.Gold + monster_gold
						stats.Exp_Level()
						loop = true
					} else {
						fmt.Println("\n\nYou died but people gave you 10 gold for honor.\n\n")
						market.Gold = market.Gold + 10
						loop = true
					}
				}
			case decision == 2:
				Monsters()
				loop = true
			case decision == 3:
				loop = true
			default:
				fmt.Println("Please choose one od the options.")
			}
		}
	}
}

type monsters struct {
	name   string
	hp     int
	damage int
	gold   int
	exp    int
	armour int
}
