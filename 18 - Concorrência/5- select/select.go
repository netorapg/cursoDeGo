package main

import "time"

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		select {
		case mensagemCanal1 := <-canal1:
			println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			println(mensagemCanal2)
		}

		/*mensagemCanal1 := <-canal1
		println(mensagemCanal1)
		mensagemCanal2 := <-canal2
		println(mensagemCanal2)*/
	}

}
