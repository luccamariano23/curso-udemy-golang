package main

import "fmt"

func notaParaConceito(nota float64) string {

	switch { //case true
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 7 && nota < 9:
		return "B"
	default:
		return "C"
	}

}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(7.4))
	fmt.Println(notaParaConceito(3.5))
}
