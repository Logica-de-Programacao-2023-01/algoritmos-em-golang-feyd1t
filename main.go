package main

import "fmt"

func main() {
	var idade = 24
	var peso = 70
	var name string
	fmt.Println("Qual é o seu nome?")
	fmt.Println("Qual è sua idade?")
	fmt.Println("Qual é o seu peso?")
	fmt.Scan(&name)
	fmt.Println("Seu nome é:", name)
	fmt.Println("Sua idade é:", idade)
	fmt.Println("Seu peso é:", peso)

}
