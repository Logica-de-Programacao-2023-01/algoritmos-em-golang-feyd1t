package main

import "fmt"

func main() {
	var valorProduto float64
	fmt.Println("Qual valor do produto?")
	fmt.Scan(&valorProduto)

	desconto := valorProduto * 0.1
	resultado := valorProduto - desconto
	fmt.Println("O valor do é de ", resultado)

}
