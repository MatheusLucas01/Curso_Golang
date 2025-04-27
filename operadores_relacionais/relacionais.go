package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings: ", "banana" == "banana") // igual
	fmt.Println("!=", 3 != 2)                      // diferente
	fmt.Println("<", 3 < 2)                        // menor
	fmt.Println(">", 3 > 2)                        // maior

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2)) // igual

	validaNome()
}

func validaNome() {
	var nome string
	var senha string

	for {
		fmt.Printf("Digite seu nome: ")
		fmt.Scanln(&nome)

		fmt.Printf("Digite sua senha: ")
		fmt.Scanln(&senha)

		if nome != "Matheus" {
			fmt.Println("Nome errado, tente novamente!")
		} else if senha != "123" {
			fmt.Println("Senha incorreta, tente novamente!")
		} else {
			fmt.Println("Acessando...")
			break // Sai do loop quando tudo estiver correto
		}
	}
}
