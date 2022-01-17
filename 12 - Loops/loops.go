package main

import (
	"fmt"
)

func main() {

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J")
	}

	nomes := []string{"Joao", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println(usuario)

}
