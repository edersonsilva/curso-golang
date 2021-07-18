package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Ederson", "Pega na minha rola", 10)
	// go fale("Maria", "Eu...", 500)
	// go fale("Jõao", "Opa...", 500)
	go fale("Maria", "Entendii!!!", 10)
	fale("Jõao", "Parabéns", 5)
}
