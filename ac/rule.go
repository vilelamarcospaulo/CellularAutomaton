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

func (r *Rule) Run(grid *Grid) {
	rowLength := len(grid.Cells)
	ruleIndex := 0

	x := grid.Cells
	if r.Radius == 2 {
		a := x[rowLength-2] // grid.Cells[147]
		b := x[rowLength-1] //grid.Cells[148]
		c := x[0]           //grid.Cells[0]
		d := x[1]           // grid.Cells[1]
		e := x[2]           // grid.Cells[2]

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

			grid.Buffer[i] = r.Items[ruleIndex]

			a = b
			b = c
			c = d
			d = e

			posE := i + 3
			if posE >= rowLength {
				posE = posE - rowLength
			}

			e = x[posE]
		}
	} else if r.Radius == 3 {
		a := grid.Cells[rowLength-3]
		b := grid.Cells[rowLength-2]
		c := grid.Cells[rowLength-1]
		d := grid.Cells[0]
		e := grid.Cells[1]
		f := grid.Cells[2]
		g := grid.Cells[3]

		for i := 0; i < rowLength; i++ {
			if a == 0 { //0
				if b == 0 { // 00
					if c == 0 { // 000
						if d == 0 { // 0000
							if e == 0 { // 00000
								if f == 0 { // 000000
									if g == 0 { // 000000
										ruleIndex = 0
									} else { // 000001
										ruleIndex = 1
									}
								} else { // 000001
									if g == 0 { // 000010
										ruleIndex = 2
									} else { // 000011
										ruleIndex = 3
									}
								}
							} else { // 00001
								if f == 0 { // 000010
									if g == 0 { // 000100
										ruleIndex = 4
									} else { // 000001
										ruleIndex = 5
									}
								} else { // 000011
									if g == 0 { // 0000110
										ruleIndex = 6
									} else { // 0000111
										ruleIndex = 7
									}
								}
							}
						} else { // 0001
							if e == 0 { // 00010
								if f == 0 { // 000100
									if g == 0 { // 0001000
										ruleIndex = 8
									} else { // 0001001
										ruleIndex = 9
									}
								} else { // 000101
									if g == 0 { // 0001010
										ruleIndex = 10
									} else { // 0001010
										ruleIndex = 11
									}
								}
							} else { // 00011
								if f == 0 { // 000110
									if g == 0 { // 0001100
										ruleIndex = 12
									} else { // 0001101
										ruleIndex = 13
									}
								} else { // 000111
									if g == 0 { // 0001110
										ruleIndex = 14
									} else { // 0001111
										ruleIndex = 15
									}
								}
							}
						}
					} else { // 001
						if d == 0 { // 0010

							if e == 0 { // 00100
								if f == 0 { // 001000
									if g == 0 { // 0010000
										ruleIndex = 16
									} else { // 00001001
										ruleIndex = 17
									}
								} else { // 001001
									if g == 0 { // 0010010
										ruleIndex = 18
									} else { // 0010011
										ruleIndex = 19
									}
								}
							} else { // 00101
								if f == 0 { // 001010
									if g == 0 { // 0010100
										ruleIndex = 20
									} else { // 001011
										ruleIndex = 21
									}
								} else { // 001011
									if g == 0 { // 0010110
										ruleIndex = 22
									} else { // 0010111
										ruleIndex = 23
									}
								}
							}
						} else { // 0011
							if e == 0 { // 00110
								if f == 0 { // 001100
									if g == 0 { // 0011000
										ruleIndex = 24
									} else { // 0011001
										ruleIndex = 25
									}
								} else { // 001101
									if g == 0 { // 0011010
										ruleIndex = 26
									} else { // 0011010
										ruleIndex = 27
									}
								}
							} else { // 00111
								if f == 0 { // 001110
									if g == 0 { // 0011100
										ruleIndex = 28
									} else { // 0011101
										ruleIndex = 29
									}
								} else { // 001111
									if g == 0 { // 0011110
										ruleIndex = 30
									} else { // 0011111
										ruleIndex = 31
									}
								}
							}
						}
					}
				} else { // 01
					if c == 0 { // 010
						if d == 0 { // 0100
							if e == 0 { // 01000
								if f == 0 { // 010000
									if g == 0 { // 010000
										ruleIndex = 32
									} else { // 010001
										ruleIndex = 33
									}
								} else { // 010001
									if g == 0 { // 010010
										ruleIndex = 34
									} else { // 010011
										ruleIndex = 35
									}
								}
							} else { // 01001
								if f == 0 { // 010010
									if g == 0 { // 010100
										ruleIndex = 36
									} else { // 010001
										ruleIndex = 37
									}
								} else { // 010011
									if g == 0 { // 0100110
										ruleIndex = 38
									} else { // 0100111
										ruleIndex = 39
									}
								}
							}
						} else { // 0101
							if e == 0 { // 01010
								if f == 0 { // 010100
									if g == 0 { // 0101000
										ruleIndex = 40
									} else { // 0101001
										ruleIndex = 41
									}
								} else { // 010101
									if g == 0 { // 0101010
										ruleIndex = 42
									} else { // 0101010
										ruleIndex = 43
									}
								}
							} else { // 01011
								if f == 0 { // 010110
									if g == 0 { // 0101100
										ruleIndex = 44
									} else { // 0101101
										ruleIndex = 45
									}
								} else { // 010111
									if g == 0 { // 0101110
										ruleIndex = 46
									} else { // 0101111
										ruleIndex = 47
									}
								}
							}
						}
					} else { // 011
						if d == 0 { // 0110

							if e == 0 { // 01100
								if f == 0 { // 011000
									if g == 0 { // 0110000
										ruleIndex = 48
									} else { // 01001001
										ruleIndex = 49
									}
								} else { // 011001
									if g == 0 { // 0110010
										ruleIndex = 50
									} else { // 0110011
										ruleIndex = 51
									}
								}
							} else { // 01101
								if f == 0 { // 011010
									if g == 0 { // 0110100
										ruleIndex = 52
									} else { // 011011
										ruleIndex = 53
									}
								} else { // 011011
									if g == 0 { // 0110110
										ruleIndex = 54
									} else { // 0110111
										ruleIndex = 55
									}
								}
							}
						} else { // 0111
							if e == 0 { // 01110
								if f == 0 { // 011100
									if g == 0 { // 0111000
										ruleIndex = 56
									} else { // 0111001
										ruleIndex = 57
									}
								} else { // 011101
									if g == 0 { // 0111010
										ruleIndex = 58
									} else { // 0111010
										ruleIndex = 59
									}
								}
							} else { // 01111
								if f == 0 { // 011110
									if g == 0 { // 0111100
										ruleIndex = 60
									} else { // 0111101
										ruleIndex = 61
									}
								} else { // 011111
									if g == 0 { // 0111110
										ruleIndex = 62
									} else { // 0111111
										ruleIndex = 63
									}
								}
							}
						}
					}
				}
			} else { // 1
				if b == 0 { // 10
					if c == 0 { // 100
						if d == 0 { // 1000
							if e == 0 { // 10000
								if f == 0 { // 100000
									if g == 0 { // 100000
										ruleIndex = 64
									} else { // 100001
										ruleIndex = 65
									}
								} else { // 100001
									if g == 0 { // 100010
										ruleIndex = 66
									} else { // 100011
										ruleIndex = 67
									}
								}
							} else { // 10001
								if f == 0 { // 100010
									if g == 0 { // 100100
										ruleIndex = 68
									} else { // 100001
										ruleIndex = 69
									}
								} else { // 100011
									if g == 0 { // 1000110
										ruleIndex = 70
									} else { // 1000111
										ruleIndex = 71
									}
								}
							}
						} else { // 1001
							if e == 0 { // 10010
								if f == 0 { // 100100
									if g == 0 { // 1001000
										ruleIndex = 72
									} else { // 1001001
										ruleIndex = 73
									}
								} else { // 100101
									if g == 0 { // 1001010
										ruleIndex = 74
									} else { // 1001010
										ruleIndex = 75
									}
								}
							} else { // 10011
								if f == 0 { // 100110
									if g == 0 { // 1001100
										ruleIndex = 76
									} else { // 1001101
										ruleIndex = 77
									}
								} else { // 100111
									if g == 0 { // 1001110
										ruleIndex = 78
									} else { // 1001111
										ruleIndex = 79
									}
								}
							}
						}
					} else { // 101
						if d == 0 { // 1010

							if e == 0 { // 10100
								if f == 0 { // 101000
									if g == 0 { // 1010000
										ruleIndex = 80
									} else { // 10001001
										ruleIndex = 81
									}
								} else { // 101001
									if g == 0 { // 1010010
										ruleIndex = 82
									} else { // 1010011
										ruleIndex = 83
									}
								}
							} else { // 10101
								if f == 0 { // 101010
									if g == 0 { // 1010100
										ruleIndex = 84
									} else { // 101011
										ruleIndex = 85
									}
								} else { // 101011
									if g == 0 { // 1010110
										ruleIndex = 86
									} else { // 1010111
										ruleIndex = 87
									}
								}
							}
						} else { // 1011
							if e == 0 { // 10110
								if f == 0 { // 101100
									if g == 0 { // 1011000
										ruleIndex = 88
									} else { // 1011001
										ruleIndex = 89
									}
								} else { // 101101
									if g == 0 { // 1011010
										ruleIndex = 90
									} else { // 1011010
										ruleIndex = 91
									}
								}
							} else { // 10111
								if f == 0 { // 101110
									if g == 0 { // 1011100
										ruleIndex = 92
									} else { // 1011101
										ruleIndex = 93
									}
								} else { // 101111
									if g == 0 { // 1011110
										ruleIndex = 94
									} else { // 1011111
										ruleIndex = 95
									}
								}
							}
						}
					}
				} else { // 11
					if c == 0 { // 110
						if d == 0 { // 1100
							if e == 0 { // 11000
								if f == 0 { // 110000
									if g == 0 { // 110000
										ruleIndex = 96
									} else { // 110001
										ruleIndex = 97
									}
								} else { // 110001
									if g == 0 { // 110010
										ruleIndex = 98
									} else { // 110011
										ruleIndex = 99
									}
								}
							} else { // 11001
								if f == 0 { // 110010
									if g == 0 { // 110100
										ruleIndex = 100
									} else { // 110001
										ruleIndex = 101
									}
								} else { // 110011
									if g == 0 { // 1100110
										ruleIndex = 102
									} else { // 1100111
										ruleIndex = 103
									}
								}
							}
						} else { // 1101
							if e == 0 { // 11010
								if f == 0 { // 110100
									if g == 0 { // 1101000
										ruleIndex = 104
									} else { // 1101001
										ruleIndex = 105
									}
								} else { // 110101
									if g == 0 { // 1101010
										ruleIndex = 106
									} else { // 1101010
										ruleIndex = 107
									}
								}
							} else { // 11011
								if f == 0 { // 110110
									if g == 0 { // 1101100
										ruleIndex = 108
									} else { // 1101101
										ruleIndex = 109
									}
								} else { // 110111
									if g == 0 { // 1101110
										ruleIndex = 110
									} else { // 1101111
										ruleIndex = 111
									}
								}
							}
						}
					} else { // 111
						if d == 0 { // 1110
							if e == 0 { // 11100
								if f == 0 { // 111000
									if g == 0 { // 1110000
										ruleIndex = 112
									} else { // 11001001
										ruleIndex = 113
									}
								} else { // 111001
									if g == 0 { // 1110010
										ruleIndex = 114
									} else { // 1110011
										ruleIndex = 115
									}
								}
							} else { // 11101
								if f == 0 { // 111010
									if g == 0 { // 1110100
										ruleIndex = 116
									} else { // 111011
										ruleIndex = 117
									}
								} else { // 111011
									if g == 0 { // 1110110
										ruleIndex = 118
									} else { // 1110111
										ruleIndex = 119
									}
								}
							}
						} else { // 1111
							if e == 0 { // 11110
								if f == 0 { // 111100
									if g == 0 { // 1111000
										ruleIndex = 120
									} else { // 1111001
										ruleIndex = 121
									}
								} else { // 111101
									if g == 0 { // 1111010
										ruleIndex = 122
									} else { // 1111010
										ruleIndex = 123
									}
								}
							} else { // 11111
								if f == 0 { // 111110
									if g == 0 { // 1111100
										ruleIndex = 124
									} else { // 1111101
										ruleIndex = 125
									}
								} else { // 111111
									if g == 0 { // 1111110
										ruleIndex = 126
									} else { // 1111111
										ruleIndex = 127
									}
								}
							}
						}
					}
				}
			}

			grid.Buffer[i] = r.Items[ruleIndex]

			a = b
			b = c
			c = d
			d = e
			e = f
			f = g

			posG := i + 4
			if posG >= rowLength {
				posG = posG - rowLength
			}

			g = x[posG]
		}
	} else {
		length := 2*r.Radius + 1
		word := make([]byte, length)

		// carrega a primeira palavra (centro 0)
		for i := 0; i < length; i++ {
			word[i] = grid.Cells[calcIndex(0, i-r.Radius, rowLength)]
		}
		ruleIndex = toInteger(word)
		grid.Buffer[0] = r.Items[ruleIndex]

		pow := float64(length - 1)
		subtract := int(math.Pow(2, pow))

		for i := 1; i < rowLength; i++ {
			if word[0] == 1 {
				ruleIndex -= subtract
			}

			ruleIndex *= 2
			incoming := grid.Cells[calcIndex(i, -r.Radius, rowLength)]
			if incoming == 1 {
				ruleIndex++
			}

			grid.Buffer[i] = r.Items[ruleIndex]
			unshift(word, incoming)
		}
	}

	grid.Cells, grid.PCells, grid.PPCells, grid.Buffer = grid.Buffer, grid.Cells, grid.PCells, grid.PPCells
}
