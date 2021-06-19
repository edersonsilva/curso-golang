package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Ederson Silva":   10000.00,
		"Juarez SOares":   5000.00,
		"Tizuka Yamazaki": 3000.00,
	}

	funcsESalarios["Jaspion"] = 1000.00

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
