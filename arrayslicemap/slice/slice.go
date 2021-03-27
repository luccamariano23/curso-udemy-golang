package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)

	//Slice não é um array. Ele é um pedaço de um array (sub array)
	a2 := [5]int{1, 2, 3, 4, 5} //array
	s2 := a2[1:3]               //slice 1 ao 3, não incluindo o 3

	fmt.Println(s2)
}
