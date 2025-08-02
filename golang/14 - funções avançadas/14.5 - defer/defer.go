package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}
func statusAluno (notas1, notas2 float32) bool {
	fmt.Println("Entrando na função para verificar se o aluno está aprovado!")
	defer fmt.Println("Media calculada. Resultado será retornado!")
	
	media := (notas1) + (notas2) / 2

	if media >= 6 {
		return true
	}
	return false
}
func main() {


	fmt.Println(statusAluno(9, 11))
 }