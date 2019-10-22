package main

import "fmt"

func main()	{

	var celsius float64
	var fahrenheit float64


	fmt.Printf("Digite a temperatura em Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)

	celsius = ((fahrenheit - 32)* 5/9)

	fmt.Println("A temperatura em Celsius equivale a: ", celsius)
}
