package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str := scanner.Text()

	var carac1, carac2 string
	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scan(&carac1)
	fmt.Print("Digite o novo caractere: ")
	fmt.Scan(&carac2)

	novaStr := strings.ReplaceAll(str, carac1, carac2)

	fmt.Printf("String resultante: %s\n", novaStr)

}
