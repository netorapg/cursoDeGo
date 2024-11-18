package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	div := 1 / 2
	multiplicacao := 1 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, div, multiplicacao, restoDaDivisao)

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Operadores Lógicos
	fmt.Println("------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)

	//Operadores unários
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)

}
