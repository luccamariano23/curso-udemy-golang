package main

import "fmt"

func main() {
	//atribuição padrão
	var b byte = 3
	fmt.Println(b)

	i := 3 //determina tipagem da variável e atribui valor a mesma

	a, b := true, 2 //determina tipagem das variáveis e atribui valor as mesmas
	if a {
		fmt.Println(a, b)
	}

	i += 3 //i = i + 3 (Serve para +, -, /, *, %)
	fmt.Println(i)
}
