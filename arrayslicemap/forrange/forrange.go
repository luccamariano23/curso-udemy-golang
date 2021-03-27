package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador vai contar os elementos

	//for range é semelhante ao foreach
	for i, numero := range numeros {
		fmt.Println(i, numero)
	}
}
