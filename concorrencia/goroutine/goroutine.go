package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d) \n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("João", "Olá", 5)
	// fale("Maria", "Eai", 1)

	// go fale("João", "Ei...", 500)
	// go fale("James", "Fala...", 500)

	go fale("João", "Ei...", 10)
	fale("James", "Fala...", 5)
}
