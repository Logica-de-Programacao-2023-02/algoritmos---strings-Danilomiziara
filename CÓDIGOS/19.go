package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str := scanner.Text()

	// converte a string em um slice de bytes
	bytes := []byte(str)

	// inverte o slice de bytes
	for i := 0; i < len(bytes)/2; i++ {
		j := len(bytes) - i - 1
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	// converte o slice de bytes de volta para string
	invertida := string(bytes)

	fmt.Printf("A string invertida é: %s\n", invertida)

}
