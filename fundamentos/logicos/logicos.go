package main

import "fmt"

func compras(tab1, tab2 bool) (bool, bool, bool) {
	comprarTV50 := tab1 && tab2
	comprarTV32 := tab1 != tab2
	comprarSorvete := tab1 || tab2

	return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("TV50: %t, TV32:, %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
