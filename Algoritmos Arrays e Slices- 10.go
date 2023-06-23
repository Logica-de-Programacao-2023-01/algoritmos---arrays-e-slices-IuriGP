package main

import (
	"fmt"
	"math"
)

func main() {
	slice := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	menorNum := math.MaxInt64
	maiorNum := math.MinInt64

	for _, valor := range slice {
		if valor < menorNum {
			menorNum = valor
		}
		if valor > maiorNum {
			maiorNum = valor
		}
	}
	fmt.Println("Maior valor: ", maiorNum, "Menor valor: ", menorNum)
}
