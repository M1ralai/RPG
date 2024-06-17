package stats

import (
	"fmt"
)

var Level, Exp int

func Exp_Level() {
	var loop bool
	var decision, remain, cap int
	cap = 100 + (Level * 200)
	remain = cap - Exp
	if Exp < cap {
		fmt.Printf("\n\nYou have %v experiende and you need to earn %v experience top level up\n\n", Exp, remain)
		fmt.Println("1.Main menu")
	} else {
		Level += 1
		fmt.Printf("\n\nCongrats you leveled up and you %v level now\n\n", Level)
		Exp = Exp - cap
	}
	loop = false

	for loop == false {
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			loop = true
		default:
			fmt.Println("Please choose one of the options.")
		}
	}
}
