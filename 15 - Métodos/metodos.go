package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Fulano", 30}
	usuario1.salvar()

	usuario2 := usuario{"Ciclano", 25}
	usuario2.salvar()
	maiordeidade := usuario2.maiorDeIdade()
	fmt.Println(maiordeidade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
