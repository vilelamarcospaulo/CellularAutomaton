package ac

type Grid struct {
	PPCells **([]byte)
	PCells  **([]byte)
	Cells   **([]byte)
	T       *int
}

func (g Grid) nextStep(r Rule) {
	*(g.PPCells) = *g.PCells
	*(g.PCells) = *g.Cells
	result := r.Run(*(*g.Cells))
	*(g.Cells) = &result
	*(g.T)++
}
