package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	numeros := [5]int{x, 2, 3}
	fmt.Print(numeros)

}
