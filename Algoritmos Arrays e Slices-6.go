package main

import "fmt"

func main() {
	numeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var x int
	fmt.Scan(&x)

	if x > numeros {
		print("sim")

	} else {
		fmt.Print("nao")
	}

}
