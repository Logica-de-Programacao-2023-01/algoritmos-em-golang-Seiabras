package main

import "fmt"

func main() {

	var x int
	var y int
	var z int
	var a int

	fmt.Print("Me diz um número? ")
	fmt.Scan(&x)
	fmt.Print("Me diz outro número: ")
	fmt.Scan(&y)
	fmt.Print("Se quiser, me diz outro número (se não escreva a): ")
	fmt.Scan(&z)

	if z == a {
		if x > y {
			fmt.Println("o primeiro é maior que o segundo")
		} else if x < y {
			fmt.Println("o primeiro é menor que o segundo")
		} else {
			fmt.Println("o primeiro e o segundo são iguais")
		}
	} else {
		if x < y && x < z && x != y && x != z {
			fmt.Println(x, ", este é o menor valor entre os três")
		} else if y < x && y < z && y != x && y != z {
			fmt.Println(y, ", este é o menor valor entre os três")
		} else if z < x && z < y && z != x && z != y {
			fmt.Println(z, ", este é o menor valor entre os três")
		} else if x < z && x == y && x != z {
			fmt.Println(x, y, ", estes números são iguais")
		} else if x < y && x == z && x != y {
			fmt.Println(x, z, ", estes números são iguais")
		} else if y < x && y == z && x != y {
			fmt.Println(y, z, ", estes números são iguais")
		} else if y == x && y == z {
			fmt.Println(y, z, x, ", todos esses números são iguais")
		}

	}
}
