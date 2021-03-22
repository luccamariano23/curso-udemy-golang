package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else {
		return "C"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(7.4))
	fmt.Println(notaParaConceito(3.5))
}
