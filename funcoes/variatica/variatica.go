package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.00
	for _, i := range numeros {
		total += i
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Media: %.2f", media(4.5, 8.8, 9.8, 6.4))
}
