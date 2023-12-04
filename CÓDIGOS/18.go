package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str := scanner.Text()

	str = strings.ReplaceAll(str, " ", "") // Remove os espaços em branco da string

	// Verifica se cada caractere da string é um número
	for _, c := range str {
		if !unicode.IsDigit(c) {
			fmt.Println("A string não contém somente números.")
			return
		}
	}

	fmt.Println("A string contém somente números(Fora o caractere Espaco).")
}
