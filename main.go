package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	teste := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes == teste)

	for i := 0; i < len(primes); i++ {
		primes[i] = i
	}

	fmt.Println(primes == teste)
}
