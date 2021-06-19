package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings =>", "banana" == "banana")
	fmt.Println("!=", 2 != 3)
	fmt.Println("> Maior: ", 3 > 2)
	fmt.Println("< Mneor: ", 3 < 2)
	fmt.Println(">= Maior igual:", 3 >= 2)
	fmt.Println("<= Menor igual:", 3 <= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Eder"}
	p2 := Pessoa{"Eder"}

	fmt.Println("Pessoas:", p1 == p2)

}
