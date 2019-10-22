//uso de array
package main

import "fmt"

func main() {
	//inicializa uma array de tamanho 5
	x := [5]float64{
		98,
		56,
		78,
		32,
		75,
	}
	var total float64 = 0
	/*
		for i := 0; i < len(x); i++ {
			fmt.Println(x[i])
			total += x[i]
		}
	*/
	for _, value := range x {
		total += value
	}

	//len() retorna um numero inteiro
	//float64(len()) converte o inteiro para o tipo de total
	fmt.Println(total / float64(len(x)))
}
