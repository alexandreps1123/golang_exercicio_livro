//Programa que retorna numeros entre 1 e 100
//que sao divisiveis por 4
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
