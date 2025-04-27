package main

import (
	"fmt"
	m "math" // pode usar alguma letra para informar a importação
)

func main() {
	const PI float64 = 3.1415 // Constante de PI, e o tipo da constante
	var raio = 3.2            // tipo (float64) inferido pelo compilador

	//forma reduzida de declarar variaveis
	area := PI * m.Pow(raio, 2)
	fmt.Printf("A area do circulo é: %.2f\n", area) // Go não suporta f-strings como Python, mas fmt.Printf é equivalente.

	const (
		a = 1
		b = 2
	)
	fmt.Println(a + b)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// Em uma única linha, pode ser atribuida várias variáveis
	var e, f bool = true, false
	fmt.Println(" Um é verdadeiro, outro falso\n", e, f)

	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)

}
