package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) { // as tres condições irão retornar valores booleanos
	comprarTv50 := trab1 && trab2    // true -> e, operador binario pois tem dois operandos
	comprarTv32 := trab1 != trab2    // ou exclusivo -> false
	comprarSoverte := trab1 || trab2 // true -> ou

	return comprarTv50, comprarTv32, comprarSoverte

}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
} // negação de sorvete, operador unário: opera somente em um operando
