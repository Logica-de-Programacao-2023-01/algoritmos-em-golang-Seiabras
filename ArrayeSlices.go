package main

//Crie um Array de inteiros com 5 elementos e imprima cada elemento em uma linha separada.
import "fmt"

func main() {

	nomes := []string{"João", "Maria", "José", "Julia", "Antonio"}

	for _, nome := range nomes {

		fmt.Println(nome)

	}
}
