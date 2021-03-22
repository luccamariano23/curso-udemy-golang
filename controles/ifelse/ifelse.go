package main

import "fmt"

func mostrarNota(nota float64) {
	//No Go é obrigado colocad {} no if, mesmo quando tem somente uma linha
	if nota >= 7 {
		fmt.Println("Aprovado com nota ", nota)
	} else {
		fmt.Println("Reprovado com nota ", nota)
	}
}

func main() {
	mostrarNota(9.0)
}
