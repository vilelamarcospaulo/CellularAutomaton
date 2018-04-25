package ac

type Grid struct {
	PPCells []byte
	PCells  []byte
	Cells   []byte
	Buffer  []byte
	T       int
}

func (g *Grid) nextStep(r Rule) {
	r.Run(g)
	g.T++
}
