package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual exercício você quer? ")
	fmt.Scan(&x)

	if x == 1 {
		for i := 1; i < 11; i++ {
			if i == 11 {
				break
			}
			fmt.Println(i)
		}
		fmt.Println("Estes são os números de 1 a 10")
	} else if x == 2 {
		for i := 0; i < 21; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		fmt.Println("Estes são os números pares de 0 a 20")
	} else if x == 3 {
		for i := 0; i < 21; i++ {
			if i%2 == 1 {
				fmt.Println(i)
			}
		}

	} else if x == 4 {
		for i := 0; i < 31; i += 3 {
			fmt.Println(i)
		}
	} else if x == 5 {
		for i := 10; i > 0; i -= 1 {
			fmt.Println(i)
		}
	}
}
