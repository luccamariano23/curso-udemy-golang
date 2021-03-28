package main

import "fmt"

func main() {
	//map dentro de um outro map
	//cada indice do map contém um outro map
	funcsPorLetra := map[string]map[string]float64{
		"A": {
			"Ana":     123.1,
			"Antonio": 456.1,
		},
		"B": {
			"Bruna":    189.2,
			"Bernardo": 1200.12,
		},
	}

	delete(funcsPorLetra, "Z")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
