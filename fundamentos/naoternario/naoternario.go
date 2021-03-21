package main

import "fmt"

func obterResultado(nota float64) string {
	if nota > 6 {
		return "Aprovado"
	}

	return "Reprovado"

	//Golang não possui operador ternário
	//Ex. nota > 6 ? "Aprovado" : "Reprovado"
}

func main() {
	fmt.Println(obterResultado(7.0))
}
