package main

import (
	"./ac"
)

func main() {
	T := 0

	cells := make([]byte, 11)
	cells[5] = 1

	items := make([]byte, 8)
	items[0] = 1
	items[1] = 0
	items[2] = 1
	items[3] = 1
	items[4] = 0
	items[5] = 1
	items[6] = 0
	items[7] = 0

	grid := ac.Grid {
		Cells: &cells,
		T: &T,
	}


	rule := ac.Rule {
		Items: items,
		Radius: 1,
	}


	ac := ac.CelullarAutomaton {
		Grid: grid,
		TMax: 10,
		Rule: rule,
	}
	
	ac.Run()
}
