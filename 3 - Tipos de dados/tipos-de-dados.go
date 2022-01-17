package main

import "fmt"

func main() {
	var numero int64 = 100000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//INT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000000000000.45
	fmt.Println(numeroReal2)

	//STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//FIM STRINGS

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

}
