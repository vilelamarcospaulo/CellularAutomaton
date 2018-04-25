package ac

import "math"

type Rule struct {
	Items  []byte
	Radius int
}

func calcIndex(i int, shift int, length int) int {
	directIndex := i + shift

	if directIndex < 0 {
		return length + directIndex
	}

	if directIndex >= length {
		return directIndex - length
	}

	return directIndex
}

func toInteger(a []byte) int {
	length := len(a)
	sum := 0
	for i := 0; i < length; i++ {
		value := 0
		if a[length-i-1] == 1 {
			value = 1
		}
		j := float64(i)
		sum += value * int(math.Pow(2, j))
	}

	return sum
}

func unshift(a []byte, b byte) {
	len := len(a)
	for i := 1; i < len; i++ {
		a[i-1] = a[i]
	}
	a[len-1] = b
}

func (r *Rule) Run(g *Grid) {
	rowLength := len(g.Cells)
	ruleIndex := 0

	x := g.Cells

	a := x[147] // g.Cells[147]
	b := x[148] //g.Cells[148]
	c := x[0]   //g.Cells[0]
	d := x[1]   // g.Cells[1]
	e := x[2]   // g.Cells[2]

	for i := 0; i < rowLength; i++ {
		if a == 0 {
			if b == 0 { // 00
				if c == 0 { // 000
					if d == 0 { // 0000
						if e == 0 { // 00000
							ruleIndex = 0
						} else { // 00001
							ruleIndex = 1
						}
					} else { // 0001
						if e == 0 { // 00010
							ruleIndex = 2
						} else { // 00011
							ruleIndex = 3
						}
					}
				} else { // 001
					if d == 0 { // 0010
						if e == 0 { // 00100
							ruleIndex = 4
						} else { // 00101
							ruleIndex = 5
						}
					} else { // 0011
						if e == 0 { // 00110
							ruleIndex = 6
						} else { // 00111
							ruleIndex = 7
						}
					}
				}
			} else { // 01
				if c == 0 { // 010
					if d == 0 { // 0100
						if e == 0 { // 01000
							ruleIndex = 8
						} else { // 01001
							ruleIndex = 9
						}
					} else { // 0101
						if e == 0 { // 01010
							ruleIndex = 10
						} else { // 01011
							ruleIndex = 11
						}
					}
				} else { // 011
					if d == 0 { // 0110
						if e == 0 { // 01100
							ruleIndex = 12
						} else { // 01101
							ruleIndex = 13
						}
					} else { // 0111
						if e == 0 { // 01110
							ruleIndex = 14
						} else { // 01111
							ruleIndex = 15
						}
					}
				}
			}
		} else {
			if b == 0 { // 10
				if c == 0 { // 100
					if d == 0 { // 1000
						if e == 0 { // 10000
							ruleIndex = 16
						} else { // 10001
							ruleIndex = 17
						}
					} else { // 1001
						if e == 0 { // 10010
							ruleIndex = 18
						} else { // 10011
							ruleIndex = 19
						}
					}
				} else { // 101
					if d == 0 { // 1010
						if e == 0 { // 10100
							ruleIndex = 20
						} else { // 10101
							ruleIndex = 21
						}
					} else { // 1011
						if e == 0 { // 10110
							ruleIndex = 22
						} else { // 10111
							ruleIndex = 23
						}
					}
				}
			} else { // 11
				if c == 0 { // 110
					if d == 0 { // 1100
						if e == 0 { // 11000
							ruleIndex = 24
						} else { // 11001
							ruleIndex = 25
						}
					} else { // 1101
						if e == 0 { // 1110
							ruleIndex = 26
						} else { // 11011
							ruleIndex = 27
						}
					}
				} else { // 111
					if d == 0 { // 1110
						if e == 0 { // 11100
							ruleIndex = 28
						} else { // 11101
							ruleIndex = 29
						}
					} else { // 1111
						if e == 0 { // 11110
							ruleIndex = 30
						} else { // 11111
							ruleIndex = 31
						}
					}
				}
			}
		}

		g.Buffer[i] = r.Items[ruleIndex]

		a = b
		b = c
		c = d
		d = x[(i+3)%149]
	}

	g.Cells, g.PCells, g.PPCells, g.Buffer = g.Buffer, g.Cells, g.PCells, g.PPCells
}
