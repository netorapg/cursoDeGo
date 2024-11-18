package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Renato"
	u.idade = 23
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua genérica", 123}

	u2 := usuario{"Carla", 21, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Tavares"}
	fmt.Println(u3)
}
