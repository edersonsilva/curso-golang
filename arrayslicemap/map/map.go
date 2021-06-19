package main

import "fmt"

func main() {
	// var aprovados map[int]sting
	// mapas devel ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[213232326] = "Pedro"
	aprovados[151215151] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[151215151])
	delete(aprovados, 151215151)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
