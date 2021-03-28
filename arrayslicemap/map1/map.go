package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1234567834] = "Maria"
	aprovados[12312312312] = "Pedro"
	aprovados[12312312] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)", nome, cpf)
	}

	//deletar campo no map
	delete(aprovados, 12312312312)
	fmt.Println(aprovados)
}
