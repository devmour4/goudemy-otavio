package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps!")

	usuario := map[string]string {
		"nome": "Kevin",
		"sobrenome": "Lima",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string {
		"nome" : {
			"primeiro" : "Kevin",
			"ultimo" : "Lima",
		},
		"curso" : {
			"nome" : "Aprenda Go",
			"regime" : "Remoto",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2, "DELETADO")

	usuario2["nome"] = map[string]string{
		"primeiro" : "Kevin",
		"ultimo" : "Lima",
	}
	fmt.Println(usuario2)
}