package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("Qual seu peso em kg?")
	fmt.Scan(&peso)
	resultado := peso * 2.2046
	fmt.Println("Seu peso em libras é", resultado)

}
