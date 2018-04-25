package main

import (
	"CellularAutomaton/ac"
	"math/rand"
)

func main() {
	items := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 0}
	// items[0] = 1
	// items[1] = 0
	// items[2] = 1
	// items[3] = 1
	// items[4] = 0
	// items[5] = 1
	// items[6] = 0
	// items[7] = 0

	cells := make([][]byte, 10)
	for i := 0; i < 2; i++ {
		cells[i] = make([]byte, 51)
	}

	for i := 0; i < 50; i++ {
		if rand.Intn(100) > 50 {
			cells[0][i] = 1
		} else {
			cells[0][i] = 0
		}

		if rand.Intn(100) > 50 {
			cells[1][i] = 1
		} else {
			cells[1][i] = 0
		}
	}

	automata := ac.Create(20, cells[0], items, 2)
	automata.Run()
	//fmt.Println(automata.Grid.PCells)

}
