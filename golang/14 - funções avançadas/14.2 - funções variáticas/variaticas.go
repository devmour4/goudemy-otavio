package main

import (
	"fmt"
)

							//esse aqui
func soma(numeros ...int) int{
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
func escrever (texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
func main() {
	totaldaSoma := soma(10, 30, 45, 60, 55)
	fmt.Println(totaldaSoma)
	escrever("Olá Kevin", 10, 15, 21, 7, 9)
}