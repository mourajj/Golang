package main

import "fmt"

func main() {

	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel := "String2"

	fmt.Println(variavel, variavel1)

	//OPERADORES RELACIONAIS - Igual a qualquer linguagem de programação  > >= == <= < !=

	//OPERADORES LÓGICOS

	//fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//OPERADORES UNÁRIOS

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

}
