package main

import "fmt"

func main() {
	var n int
	fmt.Print("Informe o tamanho dos arrays: ")
	fmt.Scanln(&n)

	array1 := make([]int, n)
	array2 := make([]int, n)

	fmt.Println("Informe os elementos do primeiro array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scanln(&array1[i])
	}
	fmt.Println("Informe os elementos do segundo array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scanln(&array2[i])
	}

	array3 := make([]int, n)
	for i := 0; i < n; i++ {
		array3[i] = array1[i] + array2[i]
	}
	fmt.Println("Terceiro array (soma):", array3)
}

