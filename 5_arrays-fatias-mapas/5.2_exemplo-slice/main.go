//slice eh um segmento de um array
//uma array contem o slice que sera usado
//para criar uma slice devemos usar o funcao make
//x := make([]float64, 5, 10)
//em que 5 eh o tamanho da slice
//e 10 o tamanho do array que contem a slice

//x :=arr[0:5] - arr[low:high]
//low indica o indice em que slice comeca
//high eh o indice em que ela termina(mas nao inclui o proprio indice)
//se omitimos arr[:] eh o mesmo que arr[0:len(arr)]

package main

import "fmt"

func main() {
	//uso da funcao append

	slice1 := []int{1, 2, 3}
	//append adiciona 4 e 5 ao final da slice1
	//e atribuimos isso ao slice 2
	//o primeiro argumento eh a slice que ja existe
	//os argumentos seguintes sao concatenados a slice1

	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	//uso da funcao copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)

	fmt.Println(slice3, slice4)

	copy(slice4, slice3)

	fmt.Println(slice3, slice4)
}
