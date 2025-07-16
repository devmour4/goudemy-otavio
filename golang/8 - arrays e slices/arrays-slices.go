package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices!")

	//criando um array de 5 posições
	var array1 [5]string
	//preenchendo as posições
	array1[0] = "Posição 1"
	fmt.Println(array1)

	//criando um array de 5 posições e "preenchendo"
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	//criando um array de forma elegante e sem especificar o tamanho
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	
	//criando um slice com 5 posições
	slice := []int{11, 12, 13, 14, 15}
	fmt.Println(slice)
	
	//incrementando mais uma posição no slice
	slice = append(slice, 16)
	
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//////////////////////////////////////
	/*
	*
	* Array Interno
	*
	*/

	array := make([]float32, 4, 5)
	fmt.Println(array)
	fmt.Println(cap(array))
	array = append(array, 6)
	fmt.Println(array)
	array = append(array, 7)
	fmt.Println(array)
	fmt.Println(cap(array))
	
	array = append(array, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println(array)
	fmt.Println(cap(array))

}