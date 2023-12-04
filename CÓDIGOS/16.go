package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a primeira frase: ")
	scanner.Scan()
	str1 := scanner.Text()

	fmt.Print("Digite a segunda frase: ")
	scanner.Scan()
	str2 := scanner.Text()

	if strings.Contains(str1, str2) {
		fmt.Println("A segunda frase é uma subfrase da primeira.")
	} else {
		fmt.Println("A segunda frase não é uma subfrase da primeira.")
	}

}
