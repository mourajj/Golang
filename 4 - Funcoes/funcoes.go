package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)

}