package main

import (
	"fmt"
)

func main() {
	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)
	soma := 0

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor para o elemento %d: ", i+1)
		fmt.Scan(&slice[i])
		soma += slice[i]
	}

	fmt.Println("O slice:", slice)
	fmt.Println("A soma dos valores:", soma)
}
