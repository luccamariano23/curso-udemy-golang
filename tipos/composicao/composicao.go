package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bwm7 struct {
}

func (b bwm7) ligarTurbo() {
	fmt.Println("Baliza")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Turbo")
}

func main() {
	var bwm7 esportivoLuxuoso = bwm7{}
	bwm7.fazerBaliza()
	bwm7.ligarTurbo()
}
