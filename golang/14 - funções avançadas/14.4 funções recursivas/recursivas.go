package main

import (
	"fmt"
)
//funções recursivas
/*1*/func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
func main() {
	fmt.Println("funções recursivas!")
	posicao := uint(10)
	fmt.Println(fibonacci(posicao))
	
}