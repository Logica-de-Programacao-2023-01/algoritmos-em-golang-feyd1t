package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual o numero?")
	fmt.Scan(&x)

	var antecessor = x - 1
	sucessor := x + 1

	fmt.Println("Seu antecessor é", antecessor,10
	"Seu sucessor é", sucessor)

}
