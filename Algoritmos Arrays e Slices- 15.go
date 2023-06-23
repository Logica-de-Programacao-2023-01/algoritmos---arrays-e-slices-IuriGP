package main

import "fmt"

func main() {
	array := [10]float64{1, 2.5, 4.9, 5.06, 7, 8, 9, 2, 0.35}
	Arraymenor := []float64{}

	for _, i := range array {
		if i < 5 {
			Arraymenor = append(Arraymenor, i)
		}
	}
	fmt.Println("Array com números menores que 5: ", Arraymenor)
}
