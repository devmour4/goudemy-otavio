package main

import (
	"fmt"
)

func main() {
	func(){
		fmt.Println("função anonima")
	}()
		
	func(texto string) {
		fmt.Println(texto)
	}("Passando Parametro")

	retorno := func (texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("passando coisinhas")
	fmt.Println(retorno)	
}