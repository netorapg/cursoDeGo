package main

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	println(texto)

	funcaoNova := closure()
	funcaoNova()

}
