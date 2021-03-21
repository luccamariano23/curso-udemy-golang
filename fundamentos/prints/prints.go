package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Linha")
	fmt.Println(" diferente.")

	x := 3.141516
	fmt.Printf("O valor de x é %.2f", x)
	fmt.Println("O valor de x é ", x)

	//Printf (nela é possível formatar valores)
	//Print (apenas printa na tela)
	//Println (printa quebrando linha)
}
