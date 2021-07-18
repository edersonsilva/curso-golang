package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1 //enviando dados para channel (escrita)

	<-ch // recebendo dados de um channer (leitura)

	ch <- 2

	fmt.Println(<-ch)
}
