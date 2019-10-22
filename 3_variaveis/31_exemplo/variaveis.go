package main

import "fmt"

func main()	{
	//var + nome variavel + tipo
	//var x string = "Hello, World!"
	//modo alternativo a linha 7
	var x string
	var a string = "first"
	var b string = "second"
	//declaracao de variavel mais concisa, em que nao se precisa
	//chamar a palavra reservada var e nem especificar o tipo da variavel
	//o proprio compilador Go infere o tipo da variavel
	c := 10

	//atribui a string a x e imprime x
	x = "Hello, World!"
	fmt.Println(x)

	//atribui outra string a x e imprime x
	x = "Hello"
	fmt.Println(x)

	//concatena uma string a x (x = x + "") e imprime o novo valor de x
	x += " World"
	fmt.Println(x)

	//retorna um booleano: verifica se a eh igual a b
	fmt.Println(a == b)

	fmt.Println(c)
}
