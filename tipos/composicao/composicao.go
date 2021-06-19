package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoluxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var car1 esportivoluxuoso = bmw7{}
	car1.ligarTurbo()
	car1.fazerBaliza()

	var car2 luxuoso = bmw7{}
	car2.fazerBaliza()

}
