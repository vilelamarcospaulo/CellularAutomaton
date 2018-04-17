package ac

type Grid struct {
	Cells *([]byte)
	T     *int
}

func (g Grid) nextStep(r Rule) {
	 result := r.Run(*g.Cells)
	 *(g.Cells) = result
	 *(g.T)++
}
