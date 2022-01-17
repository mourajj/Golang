package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"teste", "teste", "teste", "teste", "teste"}
	fmt.Println(array2)
	array2[4] = "posicao 6"

	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13}
	fmt.Println(slice)

	slice = append(slice, 18) // adicionar itens ao slice

	slice2 := array2[1:3]
	slice3 := slice[1:3]
	fmt.Println(slice2)
	fmt.Println(slice3)

	//Arrays Internos
	fmt.Println("--------------------")
	slice4 := make([]float32, 10, 11)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 5)
	slice4 = append(slice4, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
