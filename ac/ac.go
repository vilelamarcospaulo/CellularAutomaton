package ac

import (
	"fmt"
)

type Ac struct {
	grid Grid
	tMax int
}

func (a Ac) run() {
	for a.tMax > a.grid.t {
		a.grid.r.apply(a.grid)
		fmt.Println(a.grid.cells)
	}
}
