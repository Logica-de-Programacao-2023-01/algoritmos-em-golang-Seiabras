package main

import "fmt"

// Crie um Array de floats com 6 elementos e calcule a média dos valores armazenados no Array.
func main() {
	numeros := [6]float64{5, 5, 5, 5, 5, 5}
	var soma float64

	for _, numero := range numeros {
		soma += numero
	}

	media := soma / float64(len(numeros))
	fmt.Println(media)

}
