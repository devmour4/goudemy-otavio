package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, Go!")
	auxiliar.Escrever()
	

	err := checkmail.ValidateFormat("prod@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	
}