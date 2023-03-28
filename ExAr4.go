package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
func main() {
	var numeros []int
	var x int
	var y int
	var n int
	fmt.Println("Qual tamanho?")
	fmt.Scan(&x)
	fmt.Println("Me de um número.")
	fmt.Scan(&y)
	for {
		fmt.Println("Mais números? (se não escreva n)")
		fmt.Scan(&y)
		if y == n {
			break
		} else {
			numeros = append(numeros, y)
			fmt.Println(numeros)
		}
	}
}
