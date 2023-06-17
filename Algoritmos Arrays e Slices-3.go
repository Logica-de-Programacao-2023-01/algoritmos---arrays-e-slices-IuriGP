package main

import "fmt"

func main() {
	numeros := [4]float64{1.5, 2.5, 3.5, 1.5}
	mult := 1.0

	for _, numero := range numeros {
		mult *= numero
	}
	fmt.Println("valor final = ", mult)
}
