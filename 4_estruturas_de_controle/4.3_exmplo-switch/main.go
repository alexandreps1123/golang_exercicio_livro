package main

import "fmt"

func main() {
	var i int32

	fmt.Printf("Digite um numero: ")
	//%d informa que o numero eh do tipo inteiro base 10
	fmt.Scanf("%d", &i)

	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("Numero nao disponivel")
	}
}
