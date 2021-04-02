package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido")
	case n == 0:
		return 1, nil
	default:
		fat, _ := fatorial(n - 1)
		return n * fat, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)
}
