//Programa que imprime numeros de 1 a 100
//que sejam multiplos de 3 com a menssagem 'Fizz'
//que sejam multiplos de 5 com a menssagem 'Buzz
//e se for multiplo de ambos 'FizzBuzz
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}
