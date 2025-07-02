package main

import (
	"fmt"
)
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
func calculosMatematicos (n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
func main() {
	soma := somar (10, 20)
	fmt.Println(soma)

//declarando uma variavel do tipo função e chamando-a posteriormente.
	var f = func ()  {
		fmt.Println("Função F")
	}
	f()

	var f2 = func (txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f2 ("Esse é o texto")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(20, 10)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
