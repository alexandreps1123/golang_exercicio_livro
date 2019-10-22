package main

import "fmt"

func main()	{
	var input float64

	//escopo eh o conjunto de lugares em que se pode usar a variavel

	//imprime uma mensagem ao usuario
	fmt.Printf("Enter a number: ")

	//Scanf é um metodo de fmt que le um comando do usuario
	fmt.Scanf("%f", &input)

	output := input*2

	fmt.Println(output)
}
