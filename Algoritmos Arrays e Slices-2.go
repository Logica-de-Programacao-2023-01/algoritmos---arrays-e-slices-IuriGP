package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	numeros = append(numeros[:2], numeros[3:]...)

	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])

	}
}
