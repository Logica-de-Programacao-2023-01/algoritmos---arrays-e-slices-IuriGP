package main

import "fmt"

func main() {
	numeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var x int
	fmt.Scan(&x)

	for _, numero := range numeros {
		if x == numero {
			fmt.Println("número encontrado")
			break
		} else {
			fmt.Println("não é o número")
		}

	}
}
