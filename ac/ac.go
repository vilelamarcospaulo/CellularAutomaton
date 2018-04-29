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
	stateLength := len(state)
	stateCurrent := make([]byte, stateLength)
	for i := 0; i < len(state); i++ {
		stateCurrent[i] = state[i]
	}

	g := Grid{
		Cells:   stateCurrent,
		Buffer:  make([]byte, stateLength),
		PCells:  make([]byte, stateLength),
		PPCells: make([]byte, stateLength),
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
