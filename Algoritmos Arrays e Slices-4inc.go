package main

import "fmt"

func main() {
	var x, tamanho int
	fmt.Println("Digite o tamanho do Slice: ")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("digite o numero")
		fmt.Scan(&x)
	}
}
