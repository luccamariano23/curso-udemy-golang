package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	//verificar o tipo de valor
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "indefinido"
	}
}

func main() {
	fmt.Println(tipo(3.2))
	fmt.Println(tipo("teste"))
	fmt.Println(tipo(1))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))

}
