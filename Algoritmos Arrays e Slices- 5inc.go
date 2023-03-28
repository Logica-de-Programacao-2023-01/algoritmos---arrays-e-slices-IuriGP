package main

import "fmt"

func main() {
	var n1 int
	fmt.Scan(&n1)
	var n2 int
	fmt.Scan(&n2)
	var n3 int
	fmt.Scan(&n3)
	var n4 int
	fmt.Scan(&n4)

	var c2l1 int
	fmt.Scan(c2l1)
	var c2l2 int
	fmt.Scan(&c2l2)
	var c2l3 int
	fmt.Scan(&c2l3)
	var c2l4 int
	fmt.Scan(&c2l4)

	matriz := [2][2]int{{n1, n2}, {n3, n4}, {n5, n6}}
	fmt.Print(matriz)

	matriz_2 := [2][2]int{{c2l1, c2l2}, {c2l3, c2l4}}
	fmt.Print(matriz_2)
}
