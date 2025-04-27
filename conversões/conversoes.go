package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // Declarando e inicializando, ponto flutuante é float64
	y := 2
	fmt.Println(x / float64(y)) // deve-se converter um tipo int para float64

	// numero inteiro
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// inteiro para string
	fmt.Println("Teste" + strconv.Itoa(123))
	fmt.Println("Teste", +123)

	//~string para int
	num, _ := strconv.Atoi("123") // retorna dois valores, o _ ignora o erro que irá retornar
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

	var nome string
	var idade int
	fmt.Println("Informe seu nome e idade: ")
	fmt.Scanf("%s %d", &nome, &idade)

	fmt.Println("Os dados informados foram:", nome, idade)

}
