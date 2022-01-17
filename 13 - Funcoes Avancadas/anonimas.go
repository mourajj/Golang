package main

import "fmt"

func main() {

	retorno := func(text string) string {
		return fmt.Sprintf("Reecebido -> %s", text)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
