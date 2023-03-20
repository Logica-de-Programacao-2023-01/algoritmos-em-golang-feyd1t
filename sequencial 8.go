package main

import "fmt"

func main() {
	var numeroDias int
	fmt.Println("Quantos dias o funcionario trabalhou?")
	fmt.Scan(&numeroDias)
	var valorDiaria float64
	fmt.Println("Qual o valor da diaria do funcionario?")
	fmt.Scan(&valorDiaria)

	salario := float64(numeroDias) * valorDiaria

	fmt.Println("O salario é", salario)
}
