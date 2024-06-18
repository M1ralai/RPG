package monster_hub

import (
	"fmt"
)

func Monsters_Hp_Indicators(monster_hp, monster_fight_hp int) {
	monster_ratio := float32(monster_fight_hp) / float32(monster_hp)
	switch {
	case monster_ratio >= 0.95:
		fmt.Printf("HP %v/%v ********************", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.90:
		fmt.Printf("HP %v/%v *******************-", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.85:
		fmt.Printf("HP %v/%v ******************--", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.80:
		fmt.Printf("HP %v/%v ****************----", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.75:
		fmt.Printf("HP %v/%v ***************-----", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.70:
		fmt.Printf("HP %v/%v **************------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.65:
		fmt.Printf("HP %v/%v *************-------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.60:
		fmt.Printf("HP %v/%v ************--------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.55:
		fmt.Printf("HP %v/%v ***********---------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.50:
		fmt.Printf("HP %v/%v **********----------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.45:
		fmt.Printf("HP %v/%v *********-----------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.40:
		fmt.Printf("HP %v/%v ********------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.35:
		fmt.Printf("HP %v/%v *******-------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.30:
		fmt.Printf("HP %v/%v ******--------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.25:
		fmt.Printf("HP %v/%v *****---------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.20:
		fmt.Printf("HP %v/%v ****----------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.15:
		fmt.Printf("HP %v/%v ***-----------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.10:
		fmt.Printf("HP %v/%v **------------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.05:
		fmt.Printf("HP %v/%v *-------------------", monster_hp, monster_fight_hp)
	case monster_ratio >= 0.00:
		fmt.Printf("HP %v/%v --------------------", monster_hp, monster_fight_hp)
	}
}

func Warrior_Hp_Indicator(hp, fight_hp int) {
	ratio := float32(fight_hp) / float32(hp)
	switch {
	case ratio >= 0.95:
		fmt.Printf("HP %v/%v ********************", hp, fight_hp)
	case ratio >= 0.90:
		fmt.Printf("HP %v/%v *******************-", hp, fight_hp)
	case ratio >= 0.85:
		fmt.Printf("HP %v/%v ******************--", hp, fight_hp)
	case ratio >= 0.80:
		fmt.Printf("HP %v/%v ****************----", hp, fight_hp)
	case ratio >= 0.75:
		fmt.Printf("HP %v/%v ***************-----", hp, fight_hp)
	case ratio >= 0.70:
		fmt.Printf("HP %v/%v **************------", hp, fight_hp)
	case ratio >= 0.65:
		fmt.Printf("HP %v/%v *************-------", hp, fight_hp)
	case ratio >= 0.60:
		fmt.Printf("HP %v/%v ************--------", hp, fight_hp)
	case ratio >= 0.55:
		fmt.Printf("HP %v/%v ***********---------", hp, fight_hp)
	case ratio >= 0.50:
		fmt.Printf("HP %v/%v **********----------", hp, fight_hp)
	case ratio >= 0.45:
		fmt.Printf("HP %v/%v *********-----------", hp, fight_hp)
	case ratio >= 0.40:
		fmt.Printf("HP %v/%v ********------------", hp, fight_hp)
	case ratio >= 0.35:
		fmt.Printf("HP %v/%v *******-------------", hp, fight_hp)
	case ratio >= 0.30:
		fmt.Printf("HP %v/%v ******--------------", hp, fight_hp)
	case ratio >= 0.25:
		fmt.Printf("HP %v/%v *****---------------", hp, fight_hp)
	case ratio >= 0.20:
		fmt.Printf("HP %v/%v ****----------------", hp, fight_hp)
	case ratio >= 0.15:
		fmt.Printf("HP %v/%v ***-----------------", hp, fight_hp)
	case ratio >= 0.10:
		fmt.Printf("HP %v/%v **------------------", hp, fight_hp)
	case ratio >= 0.05:
		fmt.Printf("HP %v/%v *-------------------", hp, fight_hp)
	case ratio >= 0.00:
		fmt.Printf("HP %v/%v --------------------", hp, fight_hp)
	}
}
