package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3 // i = i + 3
	i -= 3 // i = 1 - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	x, y := 1, 2 // declarou e inicializou a variável
	fmt.Println(x, y)

	x, y = y, x // como ela já está inicializada, usa-se o operador de atribuição =
	fmt.Println(y, x)
}
