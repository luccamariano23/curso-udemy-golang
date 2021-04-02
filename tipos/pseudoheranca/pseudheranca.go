package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	//o fato de colocar 'carro' dentro de 'ferrari', faz com que a struct 'ferrari' "herde" os atributos da struct 'carro'
	carro       //campos anonimos (composição)
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 100
	f.turboLigado = true

	fmt.Println(f)
}
