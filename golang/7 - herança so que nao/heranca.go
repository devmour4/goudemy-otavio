package main

import (
	"fmt"
)
type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}
type estudante struct {
	pessoa //aqui não precisa passar um tipo pessoa além desse
	curso string
	faculdade string
}
func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 21, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia de Software", "FIAP"}
	fmt.Println(e1) //se eu colocar e1.altura exibe apenas a altura, da forma atual exibe todas as infos.
}