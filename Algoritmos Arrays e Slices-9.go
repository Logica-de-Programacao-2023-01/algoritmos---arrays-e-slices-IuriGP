package main

import "fmt"

func main() {
	slice := [6]float64{}
	var valor float64
	for i := 0; i < len(slice); i++ {
		fmt.Println("Escreva os valores de Float: ")
		fmt.Scanln(&valor)
		slice[i] = valor
	}
	fmt.Print("Resultado do Slice: ", slice)
}
