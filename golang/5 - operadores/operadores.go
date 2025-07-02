package main

import (
	"fmt"
)

func main() {
	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//FIM

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	
	fmt.Println(variavel1, variavel2)

	//FIM

	//OP Relacionais

	fmt.Println( 1 > 2)
	fmt.Println( 1 < 2)
	fmt.Println( 1 >= 2)
	fmt.Println( 1 >= 2)
	fmt.Println( 1 == 2)
	fmt.Println( 1 != 2)
	
	//FIM

	//OP Lógicos
	fmt.Println("--------**--------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//FIM

	// OP Unários
	numero := 10
	
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 20
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 10
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)

	//FIM DOS OP
}