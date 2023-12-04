package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str := scanner.Text()

	if _, err := strconv.ParseFloat(str, 64); err == nil {
		fmt.Printf("%s é um número válido em float.\n", str)
	} else {
		fmt.Printf("%s não é um número válido em float.\n", str)
	}

}
