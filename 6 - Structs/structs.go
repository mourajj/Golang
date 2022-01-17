package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	fmt.Println("Structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua tal", 10}

	usuario2 := usuario{"Jonathan", 21, enderecoExemplo}
	usuario3 := usuario{idade: 22}

	fmt.Println(usuario2)
	fmt.Println(usuario3)
}
