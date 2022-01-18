package main

import "fmt"

func init() { // Executa antes que a função main
	fmt.Println("Executando a funcao init")
}

func main() {
	fmt.Println("Função main sendo executada")
}
