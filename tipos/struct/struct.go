package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: função com receiver (receptor)
//O método "pertence" ao struct criado
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    5.35,
		desconto: 0.23,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
}
