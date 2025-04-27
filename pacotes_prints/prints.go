package main

import "fmt"

func main() {
	fmt.Print("Texto") // Imprime na mesma linha
	fmt.Print("Texto")

	fmt.Println("Texto") // Imprime pulando uma linha
	fmt.Println("Texto") // Imprime pulando uma linha

	x := 3.141516 // Variável do tipo float64
	fmt.Println("O valor de x é: ", x)

	// vai retornar a variável acima no tipo string.
	xs := fmt.Sprint(x) //Variável do tipo string
	fmt.Println("O valor de x é " + xs)

	//Print formatado
	fmt.Printf("O valor de x é %.2f", x) // para formatar com duas casas decimais

	a := 1
	b := 2.45
	c := false
	d := "Olá"
	fmt.Printf("\n %d %f %t %s", a, b, c, d)

}

// %f --> imprime do tipo float
// %d --> imprime do tipo inteiro
// %t --> Imprime do tipo boolean
// %s --> Imprime do tipo string
// %v --> Imprime todos os valores
