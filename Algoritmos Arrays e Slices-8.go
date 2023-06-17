package main

import "fmt"

func main() {
	slice := []string{"apple", "banana", "orange", "apple", "grape", "apple", "pear", "apple"}

	fmt.Println("O slice atual é: ", slice)

	var valor string
	fmt.Print("Digite um valor: ")
	fmt.Scanln(&valor)

	novoslice := make([]string, 0, len(slice))
	for _, i := range slice {
		if i != valor {
			novoslice = append(novoslice, i)
		}
	}
	fmt.Print("slice atualizado: ", novoslice)
}
