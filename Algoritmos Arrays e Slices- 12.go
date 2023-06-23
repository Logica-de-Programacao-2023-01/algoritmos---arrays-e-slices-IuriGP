package main

import "fmt"

func main() {
	Array := [5]int64{3, 6, 7, 8, 9}
	Novoarray := []int64{}

	for _, i := range Array {
		if i%3 == 0 {
			Novoarray = append(Novoarray, i)
		}
	}
	fmt.Println("Novo array ", Novoarray)
}
