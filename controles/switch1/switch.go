package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		//fallthrough faz com que o case seguinte seja executado, ou seja,
		//uma vez que a nota for 10, o programa executará o case 10 e o case 9,8
		fallthrough
	case 9, 8:
		return "A"
	default:
		return "F"
	}
}

func main() {

	fmt.Println(notaParaConceito(8.5))

}
