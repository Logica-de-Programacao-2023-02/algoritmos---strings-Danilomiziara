package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str := scanner.Text()

	Palavras := strings.Fields(str)
	fmt.Printf("Essa frase tem %d palavras", len(Palavras))

}
