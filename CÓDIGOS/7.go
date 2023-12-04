package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Print("Digite uma frase: ")
	fmt.Scan(&str)

	TemNumero := false

	for _, char := range str {
		if unicode.IsDigit(char) {
			TemNumero = true
			break
		}
	}

	if TemNumero {
		fmt.Printf("A frase '%s' contém pelo menos um número.\n", str)
	} else {
		fmt.Printf("A frase '%s' não contém nenhum número.\n", str)
	}

}
