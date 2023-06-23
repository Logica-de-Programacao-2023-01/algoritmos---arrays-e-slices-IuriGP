package main

import "fmt"

func main() {
	var x int64

	array := [7]int64{}

	fmt.Println("Indique qual será o primeiro e o último número do array")
	fmt.Scanln(&x)

	array[0] = x
	array[len(array)-1] = x

	fmt.Println(array)
}
