package main

import "math"

// Iniciando com letra maiuscula é Publico (visivel dentro e fora ddo pacote)
// Iniciando com letra minuscula é Privado (visivel apenas dentro do pacote)

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.x - b.y)
	return
}

//Distancia e responsavel por calcular a distancia linear entre 2 pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
