package ac

type Grid struct {
	cells []bool
	t     int
	r     Rule
}

func (g Grid) nextStep() {
	g.r.apply(g)
}
