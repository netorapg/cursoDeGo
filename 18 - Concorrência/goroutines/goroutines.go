// Concorrência != Paralelismo
// Concorrência é a capacidade de lidar com muitas coisas ao mesmo tempo
// Paralelismo é a capacidade de executar várias coisas ao mesmo tempo

package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!") //goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
