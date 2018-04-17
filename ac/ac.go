package ac

import (
	"fmt"
)

type CelullarAutomaton struct {
	Grid Grid
	TMax int
	Rule Rule
}

func (a CelullarAutomaton) Run() {
	for a.TMax > *a.Grid.T {
		a.Grid.nextStep(a.Rule)
		fmt.Println(a.Grid.Cells)
	}
}
