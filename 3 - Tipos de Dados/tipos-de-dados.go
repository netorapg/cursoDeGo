package main

import (
	"errors"
	"fmt"
)

func main() {
	//inteiros
	//int8 int16 int32 int64
	//uint8 uint16 uint32 uint64
	//int = int32 ou int64
	//uint = uint32 ou uint64
	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 1000
	fmt.Println(numero3)

	//Byte = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	numero5 := 1000000000000000000
	fmt.Println(numero5)

	//números reais
	//float32 float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.43
	fmt.Println(numeroReal2)

	numeroReal3 := 1230000000.43
	fmt.Println(numeroReal3)

	//strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	//Char é usado para verificar a tablea ASCII do caractere
	char := 'B'
	fmt.Println(char)

	//boelanos
	var booleano1 bool = true
	fmt.Println(booleano1)

	//erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
