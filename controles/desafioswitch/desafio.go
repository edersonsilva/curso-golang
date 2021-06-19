package main

import (
	"fmt"
)

func notaParaConceito(nota float64) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	case nota >= 1 && nota < 2:
		return "E"
	default:
		return "Nota inválida !"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.7))
	fmt.Println(notaParaConceito(8.3))
	fmt.Println(notaParaConceito(5.8))
	fmt.Println(notaParaConceito(3))
	fmt.Println(notaParaConceito(1))
	fmt.Println(notaParaConceito(0))
}
