package market

import (
	"fmt"
)

var Itemarmour int

func armour_list() {
	Body_Armour_1 := armour{"Wooden body armour", 15, 50}
	Body_Armour_2 := armour{"Copper body armour", 25, 100}
	Body_Armour_3 := armour{"Iron body armour", 35, 150}
	Body_Armour_4 := armour{"Chainmail body armour", 40, 200}

	Face_Armour_1 := armour{"Wooden helmet", 10, 50}
	Face_Armour_2 := armour{"Copper helmet", 15, 100}
	Face_Armour_3 := armour{"Iron helmet", 20, 150}
	Face_Armour_4 := armour{"Chainmail helmet", 25, 200}

	Leg_Armour_1 := armour{"Wooden leg protector", 15, 50}
	Leg_Armour_2 := armour{"Copper leg protector", 20, 100}
	Leg_Armour_3 := armour{"Iron leg protector", 25, 150}
	Leg_Armour_4 := armour{"Chainmail leg protector", 30, 200}

	fmt.Println("----Name-------------Cost---------Armour-----------")
	fmt.Println("1. ", Body_Armour_1.name, "....", Body_Armour_1.cost, "..........", Body_Armour_1.armour)
	fmt.Println("2. ", Body_Armour_2.name, "....", Body_Armour_2.cost, ".........", Body_Armour_2.armour)
	fmt.Println("3. ", Body_Armour_3.name, "......", Body_Armour_3.cost, ".........", Body_Armour_3.armour)
	fmt.Println("4. ", Body_Armour_4.name, "..", Body_Armour_4.cost, ".........", Body_Armour_4.armour)
	fmt.Println()
	fmt.Println("---------------------- HELMETS --------------------")
	fmt.Println()
	fmt.Println("----Name-------------Cost---------Armour-----------")
	fmt.Println("5. ", Face_Armour_1.name, "....", Face_Armour_1.cost, "..........", Face_Armour_1.armour)
	fmt.Println("6. ", Face_Armour_2.name, "....", Face_Armour_2.cost, ".........", Face_Armour_2.armour)
	fmt.Println("7. ", Face_Armour_3.name, "......", Face_Armour_3.cost, ".........", Face_Armour_3.armour)
	fmt.Println("8. ", Face_Armour_4.name, "..", Face_Armour_4.cost, ".........", Face_Armour_4.armour)
	fmt.Println()
	fmt.Println("-------------------- LEG ARMOURS ------------------")
	fmt.Println()
	fmt.Println("----Name-------------Cost---------Armour-----------")
	fmt.Println("9. ", Face_Armour_1.name, "....", Face_Armour_1.cost, "..........", Face_Armour_1.armour)
	fmt.Println("10. ", Face_Armour_2.name, "....", Face_Armour_2.cost, ".........", Face_Armour_2.armour)
	fmt.Println("11. ", Face_Armour_3.name, "......", Face_Armour_3.cost, ".........", Face_Armour_3.armour)
	fmt.Println("12. ", Face_Armour_4.name, "..", Face_Armour_4.cost, ".........", Face_Armour_4.armour)
	fmt.Println("13. Marketplace")
	fmt.Println("14. Main menu")

	var loop bool
	var decision, cost, armour int
	fmt.Printf("Welcome to armoursmith warrior. Here is some armours. \n Tier1: Wooden armour; 25 armour; 50 gold \n Tier2: Copper armour; 50 armour; 100 gold \n Tier3: Iron armour; 75 armour; 250 gold \n Tier4: Chainmail Armour; 90 armour; 500 Gold \n You have %v gold \n 5. Main menu\n", Gold)
	loop = false
	for !loop {
		for !loop {
			fmt.Scanln(&decision)
			switch {
			case decision == 1:
				cost = Body_Armour_1.cost
				armour = Body_Armour_1.armour
				loop = true
			case decision == 2:
				cost = Body_Armour_2.cost
				armour = Body_Armour_2.armour
				loop = true
			case decision == 3:
				cost = Body_Armour_3.cost
				armour = Body_Armour_3.armour
				loop = true
			case decision == 4:
				cost = Body_Armour_4.cost
				armour = Body_Armour_4.armour
				loop = true
			case decision == 5:
				cost = Face_Armour_1.cost
				armour = Face_Armour_1.armour
				loop = true
			case decision == 6:
				cost = Face_Armour_2.cost
				armour = Face_Armour_2.armour
				loop = true
			case decision == 7:
				cost = Face_Armour_3.cost
				armour = Face_Armour_3.armour
				loop = true
			case decision == 8:
				cost = Face_Armour_4.cost
				armour = Face_Armour_4.armour
				loop = true
			case decision == 9:
				cost = Leg_Armour_1.cost
				armour = Leg_Armour_1.armour
			case decision == 10:
				cost = Leg_Armour_2.cost
				armour = Leg_Armour_2.armour
			case decision == 11:
				cost = Leg_Armour_3.cost
				armour = Leg_Armour_3.armour
			case decision == 12:
				cost = Leg_Armour_4.cost
				armour = Leg_Armour_4.armour
			case decision == 13:
				Market_Section()
				loop = true
			case decision == 14:
				cost = 0
				loop = true
			default:
				fmt.Println("Please select one of them")
			}
		}
		if cost > 0 {
			loop = false
			if Gold >= cost {
				Itemarmour += armour
				Gold -= cost
				loop = true
			} else {
				fmt.Println("You don't have enough money.")
			}
		}
	}
}

type armour struct {
	name   string
	armour int
	cost   int
}
