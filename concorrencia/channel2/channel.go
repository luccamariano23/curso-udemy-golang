package main

import (
	"fmt"
	"time"
)

//Channel (canal) - é a forma de comunicação entre goroutines
//channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base //enviando dados para o canal

	time.Sleep(3 * time.Second)
	c <- 4 * base //enviando dados para o canal
}

func main() {
	ch := make(chan int)
	go doisTresQuatroVezes(2, ch)

	a, b := <-ch, <-ch
	fmt.Println(a, b)

	fmt.Println(<-ch)
}
