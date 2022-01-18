package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	//Não altera o valor da varíavel
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//Altera o valor da variável pois a função troca o valor no endereço de memória (ponteiro)
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) // & = Endereço de memória da váriavel
	fmt.Println(novoNumero)

}
