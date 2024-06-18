package market

import (
	"fmt"
)

var Itemdamage int

func Weapon_List() {
	var decision, cost, damage, armour int
	var loop bool
	sword_1 := one_handed_weapon{"Wooden Sword", 50, 25}
	sword_2 := one_handed_weapon{"Copper Sword", 100, 50}
	sword_3 := one_handed_weapon{"Iron Sword", 250, 75}
	sword_4 := one_handed_weapon{"Obsidian Sword", 500, 100}
	sword_5 := two_handed_weapon{"Wooden Sword", 75, 50, -5}
	sword_6 := two_handed_weapon{"Copper Sword", 150, 100, -20}
	sword_7 := two_handed_weapon{"Iron Sword", 375, 150, -40}
	sword_8 := two_handed_weapon{"Obsidian Sword", 750, 200, -50}
	fmt.Println("----Name-------------Cost---------Damage-----------")
	fmt.Println("1. ", sword_1.name, "....", sword_1.cost, "..........", sword_1.damage)
	fmt.Println("2. ", sword_2.name, "....", sword_2.cost, ".........", sword_2.damage)
	fmt.Println("3. ", sword_3.name, "......", sword_3.cost, ".........", sword_3.damage)
	fmt.Println("4. ", sword_4.name, "..", sword_4.cost, ".........", sword_4.damage)
	fmt.Println()
	fmt.Println("--------------------Two Handed-------------------")
	fmt.Println()
	fmt.Println("----Name-------------Cost---------Damage------Armour")
	fmt.Println("5. ", sword_5.name, "....", sword_5.cost, "..........", sword_5.damage, "........", sword_5.armour)
	fmt.Println("6. ", sword_6.name, "....", sword_6.cost, ".........", sword_6.damage, ".......", sword_6.armour)
	fmt.Println("7. ", sword_7.name, "......", sword_7.cost, ".........", sword_7.damage, ".......", sword_7.armour)
	fmt.Println("8. ", sword_8.name, "..", sword_8.cost, ".........", sword_8.damage, ".......", sword_8.armour)
	fmt.Println("9. Marketplace")
	fmt.Println("10. Main menu")
	for !loop {
		for !loop {
			fmt.Scanln(&decision)
			switch {
			case decision == 1:
				cost = sword_1.cost
				damage = sword_1.damage
				armour = 0
				loop = true
			case decision == 2:
				cost = sword_2.cost
				damage = sword_2.damage
				armour = 0
				loop = true
			case decision == 3:
				cost = sword_3.cost
				damage = sword_3.damage
				armour = 0
				loop = true
			case decision == 4:
				cost = sword_4.cost
				damage = sword_4.damage
				armour = 0
				loop = true
			case decision == 5:
				cost = sword_5.cost
				damage = sword_5.damage
				armour = sword_5.armour
				loop = true
			case decision == 6:
				cost = sword_6.cost
				damage = sword_6.damage
				armour = sword_6.armour
				loop = true
			case decision == 7:
				cost = sword_7.cost
				damage = sword_7.damage
				armour = sword_7.armour
				loop = true
			case decision == 8:
				cost = sword_8.cost
				damage = sword_8.damage
				armour = sword_8.armour
				loop = true
			case decision == 9:
				Market_Section()
				loop = true
			case decision == 10:
				cost = 0
				damage = 0
				loop = true
			default:
				fmt.Println("Please select one of them")
			}
		}
		if cost > 0 {
			loop = false
			if Gold >= cost {
				Itemarmour += armour
				Itemdamage = damage
				Gold -= cost
				loop = true
			} else {
				fmt.Println("You don't have enough money.")
			}
		}
	}
}

type one_handed_weapon struct {
	name   string
	cost   int
	damage int
}
type two_handed_weapon struct {
	name   string
	cost   int
	damage int
	armour int
}
