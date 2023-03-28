package main

func main() {
	numeros := [3]int{1, 2, 3}
	var soma int
	for _, numero := range numeros {
		soma += numero
		print(soma)
	}
}
