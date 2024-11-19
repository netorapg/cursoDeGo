package main

import "fmt"

var n int

func init() {
	fmt.Println("Dentro da função init")
	n = 10
}

func main() {
	fmt.Println("Dentro da função main")
	fmt.Println(n)
}
