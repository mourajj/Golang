package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    int
}

type estudante struct {
	pessoa
	curso       string
	faculdadade string
}

func main() {
	fmt.Println("heranca")

	p1 := pessoa{"Joao", "Pedro", 20, 179}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "Unisinos"}
	fmt.Println(e1)
}
