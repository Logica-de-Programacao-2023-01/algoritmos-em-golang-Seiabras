package main

import "fmt"

func main() {
	var soma, quantidade, x int

	fmt.Println("Insira um número:")
	fmt.Scan(&x)

	if x != 0 {
		soma += x
		quantidade++
	}

	for x != 0 {
		fmt.Println("Insira um número:")

		if x != 0 {
			soma += x
			quantidade++
		}
	}

	media := soma / quantidade
	fmt.Println("A media é:", media)
}
