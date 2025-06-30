package main

import (
	"fmt"
)

func main() {
	var numero int64 = -100000000
	fmt.Println(numero)
	var numero2 uint32 = 10000
	fmt.Println(numero2)
	
	//Alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)
	
	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)
	var numeroReal2 float64 = 123456.78
	fmt.Println(numeroReal2)
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	
//STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := '8'
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)
//FIM DAS STRINGS

	var booleano bool
	fmt.Println(booleano)

	var erro error 
	fmt.Println(erro)


}