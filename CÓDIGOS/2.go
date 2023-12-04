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

	str = strings.Join(strings.Split(str, " "), "")

	fmt.Printf("frase resultante: %s\n", str)

}
