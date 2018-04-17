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
		fmt.Println(**a.Grid.Cells)
	}
}

func Create(tMax int, state []byte, rule []byte, radius int) CelullarAutomaton {
	t := 0
	statePointer := &state

	g := Grid{
		Cells: &statePointer,
		T:     &t,
	}

	r := Rule{
		Items:  rule,
		Radius: radius,
	}

	ac := CelullarAutomaton{
		Grid: g,
		TMax: tMax,
		Rule: r,
	}

	return ac
}
