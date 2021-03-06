package main

import "fmt"

type item struct {
	produtoId int
	qtde      int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userId: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 3, 23.45},
			item{7, 100, 4.51},
		},
	}
	fmt.Println(pedido.valorTotal())
}
