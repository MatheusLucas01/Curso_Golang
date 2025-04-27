package main

import (
	"fmt"
	"math"
	"reflect" // Irá mostrar o tipo da variável
)

func main() {
	//Numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é: ", reflect.TypeOf(32.23))
	// Sem sinal (só valores positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//Com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64 // Maior valor possivel de um inteiro com 64 bits
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor da variavel dentro da tabela unicode é", i2)

	// números reais (float32, float64)
	var x float64 = 449.66
	fmt.Println("O valor da variavel x é", reflect.TypeOf(x))

	//boolean
	bo := true
	fmt.Println("O tipo da variável bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) // Negação do true

	// string
	s1 := "Olá"
	fmt.Println(s1 + " " + "Matheus") // Concatenando
	fmt.Println("O tamanho da string é", len(s1))

	//string com multiplas linhas
	s2 := `Olá
	Meu
	Nome
	é
	Matheus`
	fmt.Println(s2)

	// char?
	char := 'a'
	fmt.Println("O tipo de char é:", reflect.TypeOf(char))

	var (
		w = 3
		l = 5
	)

	fmt.Println("O valor da soma de a + b = ", w+l)
}
