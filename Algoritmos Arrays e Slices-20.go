package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 9, 5, 6, 7, 8}

	crescente := true

	for i := 1; i < len(slice); i++ {
		atual := slice[i]
		anterior := slice[i+1]

		if anterior >= atual {
			crescente = false
			break
		}
	}
	if crescente == true {
		fmt.Print("A sequência é crescente")
	} else {
		fmt.Print("A sequência não é crescente")
	}
}
