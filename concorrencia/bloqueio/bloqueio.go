package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Só depois que for lido!")
}

func main() {
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c)
	//dealock abaixo!!
	fmt.Println(<-c)
}
