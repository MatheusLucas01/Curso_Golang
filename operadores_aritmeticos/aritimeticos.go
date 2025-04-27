package main

import (
	"fmt"
	"math"
)

func aritmeticos() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("multiplicação=>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Módulo =>", a%b) // resto da divisão, no python seria //

	var nome string
	print("Digite seu nome: ")
	fmt.Scanln(&nome)
	fmt.Println("Seu nome é: ", nome)

	c := 3.5
	d := 5.8

	// outras operações usando Math...
	fmt.Println("maior =>", math.Max(float64(c), float64(d)))
	fmt.Println("maior =>", math.Min(c, d))
	fmt.Println("Exponenciação => ", math.Pow(c, d)) // 3 elevado a 2

}

func main() {
	aritmeticos()
}
