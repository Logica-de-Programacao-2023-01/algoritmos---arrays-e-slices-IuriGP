package main

func main() {
	numeros := [3]int{1, 2, 3}
	var mult int
	for _, numero := range numeros {
		mult *= numero
		print(mult)
	}
}
