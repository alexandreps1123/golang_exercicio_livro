package main

import "fmt"

func main ()	{
	//devolve o tamanho da string
	fmt.Println(len("Hello, World!"))
	//era pra devolver o caractere 1 em byte, mas n entendi o que acontece
	fmt.Println("Hello, World!"[5])
	//concatenacao de string
	fmt.Println("Hello," + "World")
}
