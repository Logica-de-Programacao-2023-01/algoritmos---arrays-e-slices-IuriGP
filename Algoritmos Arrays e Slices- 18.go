package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("Escolha um número inteiro qualquer: ")
	fmt.Scanln(&n)

	primos := make([]int, 0, n)
	num := 2

	for len(primos) < n {
		éPrimo := true

		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				éPrimo = false
				break
			}
		}

		if éPrimo {
			primos = append(primos, num)
		}
		num++
	}
	fmt.Println("Números primos encontrados: ", primos)
}
