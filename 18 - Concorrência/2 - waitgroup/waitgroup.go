// Concorrência != Paralelismo
// Concorrência é a capacidade de lidar com muitas coisas ao mesmo tempo
// Paralelismo é a capacidade de executar várias coisas ao mesmo tempo

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Go Routine 3")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Go Routine 4")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
