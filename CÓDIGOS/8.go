package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma frase: ")
	fmt.Scan(&str)

	inversaStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		inversaStr += string(str[i])
	}

	fmt.Printf("frase invertida: %s\n", inversaStr)

}
