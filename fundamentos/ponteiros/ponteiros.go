package main

import "fmt"

func main() {
	i := 1

	//Golang não tem aritmética de ponteiros
	var p *int = nil //nil = nulo

	p = &i //pegando endereço da váriavel na memória

	*p++

	fmt.Println("Endereço de p: ", p)
	fmt.Println("Valor de p: ", *p)
}
