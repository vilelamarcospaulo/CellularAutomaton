package ac

type CelullarAutomaton struct {
	Grid Grid
	TMax int
	Rule Rule
}

func (a *CelullarAutomaton) Run() {
	for a.TMax > a.Grid.T {
		//fmt.Println(a.Grid.Cells)
		a.Grid.nextStep(a.Rule)
	}
}

func Create(tMax int, state []byte, rule []byte, radius int) CelullarAutomaton {
	stateCurrent := make([]byte, len(state))
	for i := 0; i < len(state); i++ {
		stateCurrent[i] = state[i]
	}
	g := Grid{
		Cells:   stateCurrent,
		Buffer:  make([]byte, len(state)),
		PCells:  make([]byte, len(state)),
		PPCells: make([]byte, len(state)),
		T:       0,
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
