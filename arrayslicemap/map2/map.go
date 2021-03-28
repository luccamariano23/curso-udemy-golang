package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":  123.4,
		"Maria": 456.7,
		"Teste": 890.1,
	}

	//add no map
	funcsESalarios["Gabriel"] = 9090.23
	//delete no map
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
