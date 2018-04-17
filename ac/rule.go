package ac

import "math"

type Rule struct {
	Items []byte
	Radius int
}

func calcIndex(i int, radius int, length int) int {
	directIndex := i - radius

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
		if a[length - i - 1] == 1 {
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
		a[i - 1] = a[i]
	}
	a[len - 1] = b
}

func (r Rule) Run(row []byte) []byte {
	//tamanho da palavra, baseado no raio
	length := (2 * r.Radius) + 1
	word := make([]byte, length)

	//tamanho total da "fita"
	rowLength := len(row)
	newRow := make([]byte, rowLength)

	ruleIndex := 0

	// carrega a primeira palavra (centro 0)
	for i := 0; i < length; i++ {
		word[i] = row[calcIndex(i, r.Radius, rowLength)]
	}
	ruleIndex = toInteger(word)
	newRow[0] = r.Items[ruleIndex]
	
	//Com a primeira palavra carregada, basta apenas puxar o ultimo byte e deslocar para esquerda
	//Como deslocar para a esquerda é o mesmo que multiplcar por 2, e subtrair o bit mais significativo
	//Basta subtrailo e entao somar o bit entrando

	pow := float64(length - 1)
	subtract := int(math.Pow(2, pow))

	for i := 1; i < rowLength; i++ {
		if word[0] == 1 {
			ruleIndex -= subtract
		}
		
		ruleIndex *= 2

		incoming := row[calcIndex(i, -r.Radius, rowLength)]
		if incoming == 1 {
			ruleIndex++
		}

		newRow[i] = r.Items[ruleIndex]
		
		unshift(word, incoming)
	}

	return newRow
}