package main

import "fmt"

func main() {

	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que 0")
	} else {
		fmt.Println("numero é menos que 0")
	}
}
