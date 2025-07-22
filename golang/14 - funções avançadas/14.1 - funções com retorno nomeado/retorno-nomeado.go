package main

import (
	"fmt"
)
//basicamente a função retorna pra dentro da variavel soma e subtracao
func calculosMatematicos(n1, n2 int) (soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
func main() {
	fmt.Println("Funções com Retorno Nomeado")

	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)

}