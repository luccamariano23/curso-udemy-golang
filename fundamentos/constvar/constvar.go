package main

import (
	"fmt"
	//é possível criar nomes para blibiotecas importadas
	m "math"
)

func main() {
	const PI float64 = 3.415
	var raio = 3.2

	//forma reduzida de criar uma var (tipagem e valor :=)
	area := PI * m.Pow(raio, 2)

	//no Golang, não é possível criar uma varíavel sem utilizá-la
	//se uma varíavel for criada e não for utilizado, ocorre erro na compilação do código
	fmt.Print(area)

	//declarando mais de uma constante
	const (
		a = 1
		b = 2
	)

	fmt.Print(a, b)

	//declarando mais de uma varíavel em uma mesma linha **
	e, f := true, 2

	fmt.Print(e, f)
}
