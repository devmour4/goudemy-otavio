package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops!")
//
//	tempo := 0
//
//	for tempo < 10 {
//		tempo++
//		fmt.Println("Incrementando", tempo)
//		time.Sleep(time.Second)
//	}
//
//	fmt.Println("chegou a ", tempo)
//
//	for j := 0; j < 10; j++ {
//		fmt.Println("Incrementando j", j)
//		time.Sleep(time.Second)
//		
//	}

	nomes := [3]string{"Eduardo", "Pedro", "Jefferson"}
	
	for indice, conteudo := range nomes {
		fmt.Println(indice, conteudo)
	}
	
	nomesim := [3]string{"Joao", "Gustavo", "Bial"}
	for _, nome := range nomesim {
		fmt.Println(nome)
	}
	usuario := map[string] string {
		"nome": "Leonardo",
		"sobrenome": "Castro",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}