package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	fmt.Println("Executou!")
	c <- 1
	c <- 2
	c <- 3
	fmt.Println("Executou 2!")
	c <- 4
	fmt.Println("Executou 3!")
	c <- 5
	c <- 6
}

func main() {
	//quando o limite de um buffer de um canal é atingido, o canal bloqueia e não é possível inserir mais valores nele!
	ch := make(chan int, 3) //chan, buffer
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
