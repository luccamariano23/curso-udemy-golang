package main

import "fmt"

func multiplica(a, b int) (r int) {
	r = a * b
	return
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplica, 3, 4)
	fmt.Println(resultado)
}
