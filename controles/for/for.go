package main

import "fmt"

func main() {
	//golang tem apenas `for` como estrutura de repeticao
	//nao tem `while`, `do while`, etc.

	//`for` como `while`
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//`for`como `for`
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	//laco infinito
	//como o while(true)
	for {

	}
}
