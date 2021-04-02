package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // como os valores nomeados já foram atribuidos, não tem necessidade de colocar os valores no return
}

func main() {
	fmt.Println(trocar(2, 3))

	//x, y = trocar(2, 3)
}
