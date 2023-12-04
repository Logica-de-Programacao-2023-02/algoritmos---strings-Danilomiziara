package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1 string
	fmt.Println("Escreva uma frase:")
	fmt.Scan(&string1)
	string2 := strings.ToUpper(string1)
	fmt.Print("Sua string com todas as letras maiusculas: ", string2)

}
