package goarea

import "math"

//Pi é uma proporção numérica definida pela relação entre
//o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

//Circ é uma função que calcula a circunferência a partir do raio
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula a área do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Esta não é uma função visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
