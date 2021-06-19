package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmetica de ponteiro
	var p *int = nil

	p = &i
	*p++
	i++

	fmt.Println(p, *p, &i)
}
