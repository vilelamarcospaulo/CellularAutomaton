package ac

type Pattern struct {
	neighborhood [3]bool
	output       bool
}

type Rule struct {
	items []Pattern
}

func equals(a []bool, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func (rule Rule) apply(grid Grid) {
	//only rules with at least on item and only cells with at least 3 elements, can match with a patter of type Left Center Right
	if len(rule.items) <= 0 || len(grid.cells) < 3 {
		return
	}

	for i := 0; i < len(grid.cells); i++ {
		neighborhood := [3]bool{false, false, false}
		neighborhood[1] = grid.cells[i]

		//get last cell, or the previous
		if i == 0 {
			neighborhood[0] = grid.cells[len(grid.cells)-1]
		} else {
			neighborhood[1] = grid.cells[i-1]
		}

		//get first cell, or the next
		if i == len(grid.cells)-1 {
			neighborhood[2] = grid.cells[0]
		} else {
			neighborhood[2] = grid.cells[i+1]
		}

		for _, pattern := range rule.items {
			if pattern.neighborhood == neighborhood {
				grid.cells[i] = pattern.output
				break
			}
		}
	}

	grid.t++
}
