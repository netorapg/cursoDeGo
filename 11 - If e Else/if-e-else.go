package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}
	// Anotação: Dentro do go tem uma coisa chamada "if init", basicamente, se eu crio uma variavel dentro do condicional, ela só pertence àquele escopo
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else if numero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
