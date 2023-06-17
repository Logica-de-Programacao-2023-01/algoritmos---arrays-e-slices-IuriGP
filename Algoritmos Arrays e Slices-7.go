package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 5)

	for len(slice) < 5 {
		var num int
		fmt.Print("Digite um número inteiro: ")
		fmt.Scanln(&num)

		found := false
		for _, v := range slice {
			if v == num {
				found = true
				break
			}
		}

		if !found {
			slice = append(slice, num)
		} else {
			fmt.Println("O número já está presente no slice. Tente novamente.")
		}
	}

	fmt.Println("Slice resultante:", slice)
}
