package main

import "fmt"

//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

func main() {

	var numeros = [4]float64{1, 2, 3, 4}
	var produto float64
	produto = 1
	for _, numero := range numeros {
		produto *= numero
	}
	fmt.Println(produto)

}
