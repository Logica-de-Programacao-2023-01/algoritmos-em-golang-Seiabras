package main

import "fmt"

func main() {
	var x int
	fmt.Print("Me diz um número? ")
	fmt.Scan(&x)

	if x%2 == 1 {
		fmt.Println(x, "esse múmero é ímpar")
	} else {
		fmt.Println(x, "esse múmero é par")
	}

}
