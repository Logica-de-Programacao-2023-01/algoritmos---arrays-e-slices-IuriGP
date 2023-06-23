package main

import "fmt"

func main() {
	Matriz := [2][3]int64{{1, 2, 3}, {4, 5, 6}}
	var x, y int64
	fmt.Println("Indique a posição da linha")
	fmt.Scanln(&x)

	fmt.Println("Indique a posição da coluna")
	fmt.Scanln(&y)
	x--
	y--
	fmt.Println(Matriz[x][y])
}
