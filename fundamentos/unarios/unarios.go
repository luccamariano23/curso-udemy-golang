package main

import "fmt"

func main() {
	x := 1
	y := 2

	// Golang possui apenas postfix, ou seja,
	// possui i++ , mas não possui ++i
	x++
	y--

	fmt.Println(x, y)
}
