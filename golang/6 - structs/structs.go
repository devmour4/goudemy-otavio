package main

import (
	"fmt"
)
type usuario struct {
	nome string
	idade uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero uint8	
}
func main() {
	fmt.Println("Arquivo structs")
	
	var u usuario
	u.nome = "davi"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco {"Rua dos prazeres", 223}
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)
	
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)


}