package main

import (
	"fmt"

	"./ac"
)

func main() {
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

	automata := ac.Create(10, cells, items, 1)

	automata.Run()
	fmt.Println(cells)
}
