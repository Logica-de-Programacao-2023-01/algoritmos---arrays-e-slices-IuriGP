package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var x, y int
	fmt.Print("Qual valor de x?")
	fmt.Scan(&x)
	fmt.Print("qual valor de y?")
	fmt.Scan(&y)

	aux := slice[x]
	slice[x] = slice[y]
	slice[y] = aux
}
