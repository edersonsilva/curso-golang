package main

import "fmt"

// Nao tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(4))
}
