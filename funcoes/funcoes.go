package main

import "fmt"

func somar(a int, b int) int { // a função recebe dois valores inteiros, o nome e o tipo do parametro, e irá retornar um valor do tipo inteiro
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func welcome() {
	var nome string
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)
	fmt.Println("Olá, " + nome)

	var idade int
	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	if idade > 18 {
		fmt.Println("Você é maior de idade, pode tirar carteira!")
	} else {
		fmt.Println("Você é menor de idade")
	}
}

func main() {

	resultado := somar(3, 4)
	welcome()
	imprimir(resultado)

}
