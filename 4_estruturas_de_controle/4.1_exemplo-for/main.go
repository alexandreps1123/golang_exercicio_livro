package main

import "fmt"

func main() {

	//Em go existe apenas o laco for
	//ele pode ser usado de diversas formas

	//forma 1
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//forma 2: mais parecido com C
	//em C: for(i=10; i >= 1; i--) {}
	//descrevendo os parametros:
	//Primeiro inicializa a variavel
	//Segundo a condição de parada do laco
	//Terceiro o incremento a variavel
	for j := 10; j >= 1; j-- {
		fmt.Println(j)
	}
}
