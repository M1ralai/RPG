package main

import (
	"fmt"

	"github.com/M1ralai/RPG/market"
	"github.com/M1ralai/RPG/monster_hub"
	"github.com/M1ralai/RPG/stats"
)

func main() {
	var loop bool
	var decision int
	for !loop {
		fmt.Println("Welcome to the real world warrior choose where you wanna go \n\n 1.Marketplace \n\n 2.Monsters Hub \n\n 3.Show stats")
		stats.Stats()
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			market.Market_Section()
		case decision == 2:
			monster_hub.Monsters()
		case decision == 3:
			stats.Stats_Visualisation()
		default:
			fmt.Println("Please choose one of the options.")
		}
		loop = false
	}
}
