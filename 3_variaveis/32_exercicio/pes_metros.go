package main

import "fmt"

func main ()	{
	var pes float64
	var metros float64
	const c float64 = 0.3048

	fmt.Printf("Forneça o valor em pés: ")
	fmt.Scanf("%f", &pes)

	metros = c * pes

	fmt.Println("A medida em metros eh: ", metros)
}
