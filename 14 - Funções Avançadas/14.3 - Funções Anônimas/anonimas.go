package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("Passando Parâmetro")

	fmt.Println(retorno)

}
