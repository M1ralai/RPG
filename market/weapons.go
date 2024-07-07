package market

import (
	"fmt"
	"strconv"
	"strings"
)

var Itemdamage, decision int
var Sword_Names = [9]string{"", "Wooden Sword", "Copper Sword", "Iron Sword", "Obsidian Sword"}
var Index = [9]string{"", "----------------------One  Handed---------------------", "------Name--------------Cost-----------------Damage---", "", "", "----------------------Two  Handed---------------------", "------Name--------------Cost-----------------Damage---"}
var cost = [9]int{}
var damage = [9]int{}

func Weapon_List() {
	Gold = 500
	var loop bool
	var basecost, basedamage int
	for i := 1; i <= 8; i++ {
		if (i-1)%4 == 0 {
			fmt.Printf("%s\n%s\n", Index[i], Index[i+1])
		}
		if i <= 4 {
			basecost = 50
			basedamage = 25
		} else if i <= 8 {
			basecost = 75
			basedamage = 50
		}
		cost[i] = basecost * i
		damage[i] = basedamage * i
		fmt.Printf("%d %s %s %d %s %d %s \n", i, Sword_Names[i], strings.Repeat(".", (20-len(Sword_Names[i]))), cost[i], strings.Repeat(".", (20-len(strconv.Itoa(cost[i])))), damage[i], strings.Repeat(".", (7-len(strconv.Itoa(damage[i])))))
	}
	fmt.Println("9 Main Menu")
	for !loop {
		fmt.Scanln(&decision)
		if decision > 0 {
			if decision <= 8 {
				for i := 0; i <= 8; i++ {
					if i == decision {
						if Gold < cost[i] {
							fmt.Println("You dont have enough gold.")
						} else {
							fmt.Println("Thank you for your purchase")
							Itemdamage = damage[i]
							loop = true
						}
					}
				}
			} else if decision == 9 {
				fmt.Println("You are going main menu.")
				loop = true
			} else {
				fmt.Println("choose one of the option.")
			}
		}
	}
}
