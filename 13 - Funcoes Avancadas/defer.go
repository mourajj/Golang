package main

import "fmt"

//DEFER = Adiar a execução
func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, resultado será retornado")
	fmt.Println("Entrando na função pra verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}
