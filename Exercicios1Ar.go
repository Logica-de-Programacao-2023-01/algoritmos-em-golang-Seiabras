package main

import "fmt"

// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
func main() {

	var numeros = [3]int{1, 2, 3}
	var soma int
	for _, numero := range numeros {
		soma += numero
	}
	fmt.Println(soma)
}
