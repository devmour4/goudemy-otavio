package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Olá 1"
	variavel2 := "Olá 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Olá 3"
		variavel4 string = "Olá 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Olá 5", "Olá 6"
	fmt.Println(variavel5, variavel6)

	const constante string = "Constante"
	fmt.Println(constante)

	const (
	constante1 string = "Constantezinha"
	constante2 string = "Constantationzinha"
	)
}