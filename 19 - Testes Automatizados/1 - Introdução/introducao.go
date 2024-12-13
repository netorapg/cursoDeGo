package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia Paulista")
	fmt.Println(tipoEndereco)
}
