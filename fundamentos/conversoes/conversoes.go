package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	//converte int para float
	fmt.Println(x / float64(y))

	nota := 6.9
	//converte float para int
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//converte int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//converte string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	//converte string/num para bool
	b, _ := strconv.ParseBool("false")
	fmt.Println(b)
}
